# Rate Limiting

## What It Is

- **Rate limiting** controls how often an operation can happen (e.g., 5 requests/second)
- **ELI10:** Rate limiting is a bouncer counting how many people enter per minute -- too fast and the door stays shut
- Two patterns: **ticker-based** (fixed interval) and **token bucket** (burst-friendly)

## Why It Matters

- Every backend system needs rate limiting for APIs, DB connections, external calls
- Interviewers test whether you can implement simple rate limiting with channels

## Syntax Cheat Sheet

```go
// Go: ticker-based (fixed interval)
limiter := time.NewTicker(200 * time.Millisecond)
defer limiter.Stop()
for req := range requests {
    <-limiter.C  // wait for next tick
    handle(req)
}

// Token bucket (allows bursts)
bucket := make(chan struct{}, 3) // 3 tokens
```

```python
# Python: time.sleep for fixed interval
import time
for req in requests:
    time.sleep(0.2)  # 5 req/sec
    handle(req)
```

> **Python differs**: no channel-based timers. Use `time.sleep` for fixed
> intervals or `asyncio` rate limiters for async code.

## Tiny Example

- `main.go` -- ticker-based rate limiter + token bucket with burst capacity
- `main.py` -- time.sleep limiter + token bucket using a queue

## Common Interview Traps

- **Ticker must be stopped**: `defer limiter.Stop()` to prevent goroutine leak
- **Token bucket empty = block**: consumer waits until a token is available
- **Burst vs steady rate**: ticker = steady; token bucket = allows burst up to capacity
- **Rate limiting at wrong layer**: limit at the entry point, not deep in the call stack

## What to Say in Interviews

- "For steady rate I use a ticker; for burst-friendly I use a token bucket pattern"
- "A buffered channel can act as a simple token bucket -- fill it periodically, consume before each op"
- "In production I'd use `golang.org/x/time/rate` but the channel pattern shows understanding"

## Run It

```bash
go run ./06_concurrency/10_rate_limiting/
python ./06_concurrency/10_rate_limiting/main.py
```

## TL;DR (Interview Summary)

- **Ticker**: `time.NewTicker(d)` -- steady rate, one op per tick
- **Token bucket**: buffered channel filled periodically, consumed per request
- Token bucket allows **bursts** up to bucket capacity
- Always `defer ticker.Stop()` to prevent leaks
- In production: `golang.org/x/time/rate.Limiter`
- Python: `time.sleep(interval)` or async rate limiter libraries
