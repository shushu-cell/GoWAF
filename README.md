# GoWAF

High-performance WAF fingerprinting for **mass asset scanning** in Go (rewrite inspired by [wafw00f](https://github.com/EnableSecurity/wafw00f)).

- 🚀 Built for scanning **hundreds/thousands** of targets concurrently
- 🧾 Output **JSONL** with evidence (matched headers/status)
- 🛡️ Default **passive** mode (normal GET). No bypass.

## Install

### Go install
```bash
go install github.com/shushu-cell/GoWAF/cmd/gowaf@latest

