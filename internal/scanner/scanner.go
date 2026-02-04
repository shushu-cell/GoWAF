package scanner

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/shushu-cell/GoWAF/internal/types"
)

type Config struct {
	Workers int
	Timeout time.Duration
}

type Scanner struct {
	cfg    Config
	client *http.Client
}

func New(cfg Config) *Scanner {
	if cfg.Workers <= 0 {
		cfg.Workers = 200
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 6 * time.Second
	}
	return &Scanner{
		cfg:    cfg,
		client: newHTTPClient(cfg.Timeout),
	}
}

func (s *Scanner) ScanAll(targets []string) <-chan types.Result {
	out := make(chan types.Result, s.cfg.Workers)
	jobs := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < s.cfg.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range jobs {
				out <- s.scanOne(t)
			}
		}()
	}

	go func() {
		for _, t := range targets {
			jobs <- t
		}
		close(jobs)
		wg.Wait()
		close(out)
	}()

	return out
}

func (s *Scanner) scanOne(target string) types.Result {
	start := time.Now()
	res := types.Result{
		Target:    target,
		Timestamp: time.Now(),
	}

	url := normalizeTarget(target)

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		res.Error = err.Error()
		res.CostMS = time.Since(start).Milliseconds()
		return res
	}
	req.Header.Set("User-Agent", "gowaf/0.1 (+https://github.com/shushu-cell/GoWAF)")

	resp, err := s.client.Do(req)
	if err != nil {
		res.Error = err.Error()
		res.CostMS = time.Since(start).Milliseconds()
		return res
	}
	defer resp.Body.Close()

	// read small body only (avoid heavy operations for MVP)
	_, _ = io.ReadAll(io.LimitReader(resp.Body, 64*1024))

	res.Evidence.StatusCode = resp.StatusCode
	res.Evidence.MatchedHeaders = map[string]string{}
	if v := resp.Header.Get("Server"); v != "" {
		res.Evidence.MatchedHeaders["Server"] = v
	}
	if v := resp.Header.Get("CF-RAY"); v != "" {
		res.Evidence.MatchedHeaders["CF-RAY"] = v
		res.HasWAF = true
		res.WAF = "Cloudflare"
		res.Confidence = 0.9
	}

	res.CostMS = time.Since(start).Milliseconds()
	return res
}

func normalizeTarget(t string) string {
	t = strings.TrimSpace(t)
	if strings.HasPrefix(t, "http://") || strings.HasPrefix(t, "https://") {
		return t
	}
	return "https://" + t
}
