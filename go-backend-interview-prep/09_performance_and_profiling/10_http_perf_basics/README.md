# HTTP Performance Basics

## What It Is

- **Connection reuse**: HTTP keep-alive lets clients reuse TCP connections across requests
- **ELI10:** Reusing HTTP connections is like keeping the phone line open instead of dialing again for every sentence.
- **Client configuration**: timeouts, transport settings, and connection pooling via `http.Client`

## Why It Matters

- Creating a new TCP connection per request adds ~1-100ms of latency (DNS + TLS handshake)
- **ELI10:** Every new connection means a fresh handshake -- imagine introducing yourself to your friend before every sentence.
- Interviewers test whether you configure `http.Client` properly for production use

## Syntax Cheat Sheet

```go
// Go: reuse a configured client
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
resp, err := client.Get(url)
defer resp.Body.Close()    // MUST close to reuse connection
```

```python
# Python: urllib with timeout
import urllib.request
resp = urllib.request.urlopen(url, timeout=10)
data = resp.read()
# Note: urllib has limited connection reuse
# requests.Session() would reuse, but it's not stdlib
```

> **Go**: `http.Client` with `Transport` controls connection pooling, timeouts, keep-alive.
> **Python**: `urllib` has limited reuse; `requests.Session` (third-party) does proper pooling.

## Tiny Example

- `main.go` -- compares new-client-per-request vs reused client, with timing
- `main.py` -- demonstrates urllib with timeouts and explains connection reuse limitations

## Common Interview Traps

- **Not closing response body**: connection isn't returned to pool; leaks connections
- **New http.Client per request**: loses connection reuse; each request does fresh TCP + TLS
- **No timeouts**: a stuck server hangs your goroutine forever
- **Default transport limits**: `MaxIdleConnsPerHost=2` by default -- too low for high-throughput
- **Ignoring DNS caching**: Go doesn't cache DNS by default in some environments

## What to Say in Interviews

- "I create one http.Client at startup and reuse it -- this keeps TCP connections alive"
- "I always close resp.Body so the connection returns to the pool"
- "I set Timeout on the client and configure Transport for production connection pooling"

## Run It

```bash
go run ./09_performance_and_profiling/10_http_perf_basics/
python ./09_performance_and_profiling/10_http_perf_basics/main.py
```

## TL;DR (Interview Summary)

- Reuse `http.Client` -- one client per service, not per request
- `defer resp.Body.Close()` -- must close body to return connection to pool
- Set `Timeout` on client, not just on individual requests
- Configure `Transport`: `MaxIdleConnsPerHost`, `IdleConnTimeout`
- Default `MaxIdleConnsPerHost=2` is too low for high-throughput services
- Python: `urllib` has limited reuse; `requests.Session` is better but not stdlib
- Connection reuse saves 1-100ms per request (avoids TCP + TLS handshake)
