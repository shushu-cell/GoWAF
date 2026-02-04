package types

import (
	"encoding/json"
	"time"
)

type Evidence struct {
	MatchedHeaders map[string]string `json:"matched_headers,omitempty"`
	StatusCode     int               `json:"status_code,omitempty"`
}

type Result struct {
	Target     string    `json:"target"`
	HasWAF     bool      `json:"has_waf"`
	WAF        string    `json:"waf,omitempty"`
	Confidence float64   `json:"confidence,omitempty"`
	Evidence   Evidence  `json:"evidence,omitempty"`
	Error      string    `json:"error,omitempty"`
	CostMS     int64     `json:"cost_ms"`
	Timestamp  time.Time `json:"ts"`
}

func (r Result) JSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}
