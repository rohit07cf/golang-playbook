# DNS and Timeouts

## What It Is

- **DNS lookup**: `net.LookupHost` resolves a hostname to IP addresses
- **Timeouts**: `http.Client{Timeout: ...}` and `net.DialTimeout` prevent hanging

## Why It Matters

- Network calls without timeouts can hang forever -- a production killer
- Interviewers ask about timeout layers: DNS, connect, TLS, request, response

## Syntax Cheat Sheet

```go
// Go: DNS + timeouts
addrs, _ := net.LookupHost("example.com")
client := &http.Client{Timeout: 5 * time.Second}
conn, _ := net.DialTimeout("tcp", "host:80", 2*time.Second)
```

```python
# Python: DNS + timeouts
import socket
addrs = socket.gethostbyname_ex("example.com")
urllib.request.urlopen(url, timeout=5)
socket.create_connection(("host", 80), timeout=2)
```

> **Python differs**: `socket.gethostbyname_ex` returns (hostname, aliases, ips).
> Go's `net.LookupHost` returns a string slice of IPs directly.

## Tiny Example

- `main.go` -- DNS lookup, HTTP timeout demo (local server), dial timeout
- `main.py` -- same with socket.gethostbyname_ex, urllib timeout

## Common Interview Traps

- **Default http.Client has no timeout**: always set one explicitly
- **DNS can be slow**: in corporate networks, DNS can add seconds of latency
- **Multiple timeout layers**: connect timeout, TLS handshake, response header, body read
- **Context timeout vs client timeout**: context is per-request; client timeout is global

## What to Say in Interviews

- "I always set explicit timeouts on HTTP clients and TCP dials"
- "There are multiple timeout layers: DNS, connect, TLS, headers, body"
- "For fine-grained control I use context.WithTimeout per request"

## Run It

```bash
go run ./07_io_files_networking/10_dns_and_timeouts/
python ./07_io_files_networking/10_dns_and_timeouts/main.py
```

## TL;DR (Interview Summary)

- `net.LookupHost(host)` returns `[]string` of IP addresses
- `http.Client{Timeout: d}` -- always set; default is **no timeout**
- `net.DialTimeout("tcp", addr, d)` for connect-level timeout
- Multiple layers: DNS, connect, TLS handshake, response headers, body
- Use `context.WithTimeout` for per-request control
- Python: `socket.gethostbyname_ex`, `urlopen(timeout=N)`, `create_connection(timeout=N)`
