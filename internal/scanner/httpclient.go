package scanner

import (
	"net"
	"net/http"
	"time"
)

func newHTTPClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          2000,
		MaxConnsPerHost:       200,
		MaxIdleConnsPerHost:   200,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableCompression:    false,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}
