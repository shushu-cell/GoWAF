# GoWAF

🚀 **GoWAF** 是一个面向 **大规模资产探测 (mass asset scanning)** 的 Web 应用防火墙（WAF）识别工具，使用 Go 重写，思路参考 [wafw00f](https://github.com/EnableSecurity/wafw00f)。

GoWAF is a high-performance WAF fingerprinting tool built for **high concurrency scanning** in Go (rewrite inspired by wafw00f).

- ⚡ 高并发：适合一次扫描 100/1000+ 目标
- 🧾 结构化输出：JSONL（每行一个结果），包含证据 evidence
- 🛡️ 默认安全：passive 模式（普通 GET），不做绕过、不做攻击

---

## Features | 功能特性

- **Passive detection**：normal GET + header/status based fingerprints  
  被动识别：普通请求 + 响应头/状态码特征识别
- **JSONL output with evidence**：easy to integrate into pipelines  
  输出包含 evidence，方便对接资产平台/数据管道
- **Worker pool concurrency**：`--workers` controls concurrency  
  worker pool 并发模型，适合批量探测

---

## Install | 安装

### Option A: go install（推荐）
> 适合已安装 Go 的用户。安装后可直接使用 `gowaf` 命令。

```bash
go install github.com/shushu-cell/GoWAF/cmd/gowaf@latest
