# HTTP Client Basics

## What It Is

- Go: `net/http` package provides `http.Get`, `http.Client`, request builders
- **ELI10:** An HTTP client is your program making phone calls to other servers -- always set a timeout or you'll wait on hold forever.
- Always set a **timeout** -- the default client has **no timeout**

## Why It Matters

- Every backend service calls other services over HTTP
- **ELI10:** A backend without an HTTP client is a hermit -- it can't talk to anyone else, and modern systems are all about talking.
- Interviewers test timeout handling, headers, and response body closing

## Syntax Cheat Sheet

```go
// Go: HTTP client with timeout
client := &http.Client{Timeout: 5 * time.Second}
resp, err := client.Get(url)
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
```

```python
# Python: urllib.request with timeout
import urllib.request
resp = urllib.request.urlopen(url, timeout=5)
body = resp.read()
```

> **Python differs**: `urllib.request` is stdlib but verbose.
> `requests` (third-party) is common in production but we use stdlib here.

## Tiny Example

- `main.go` -- spins up a local test server, makes GET with timeout and custom headers
- `main.py` -- same pattern with `http.server` and `urllib.request`

## Common Interview Traps

- **Default client has no timeout**: `http.Get(url)` can hang forever
- **Must close resp.Body**: leaked connections exhaust the pool
- **Check status code**: `resp.StatusCode` is just an int -- 4xx/5xx are not errors
- **Content-Type header**: always set it for POST/PUT requests
- **Connection reuse**: closing Body promptly enables HTTP keep-alive

## What to Say in Interviews

- "I always create an http.Client with an explicit Timeout -- the default has none"
- "I defer resp.Body.Close() to prevent connection leaks"
- "For production I'd also set transport-level timeouts and connection pool limits"

## Run It

```bash
go run ./07_io_files_networking/08_http_client_basics/
python ./07_io_files_networking/08_http_client_basics/main.py
```

## TL;DR (Interview Summary)

- Always use `http.Client{Timeout: N}` -- never bare `http.Get`
- Always `defer resp.Body.Close()` after checking err
- `resp.StatusCode` is an int -- check it explicitly
- Set `Content-Type` header for POST/PUT
- Close Body promptly to enable connection reuse
- Python: `urllib.request.urlopen(url, timeout=N)`
