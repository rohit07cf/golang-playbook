# Rate Limit Basics

## What It Is

- **Rate limiting**: restricting how many requests a client can make in a time window
- **ELI10:** Rate limiting is a turnstile at the subway -- it lets people through one at a time, no matter how big the crowd.
- Common algorithms: token bucket, fixed window, sliding window

## Why It Matters

- Prevents abuse, protects backend resources, ensures fair usage
- **ELI10:** Without rate limiting, one greedy client can eat all the pizza before anyone else gets a slice.
- Interviewers ask you to implement a simple rate limiter from scratch

## Syntax Cheat Sheet

```go
// Go: simple token bucket (concurrency-safe)
type RateLimiter struct {
    mu       sync.Mutex
    tokens   int
    max      int
    interval time.Duration
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}
```

```python
# Python: simple token bucket with threading.Lock
class RateLimiter:
    def __init__(self, max_tokens, refill_interval):
        self.tokens = max_tokens
        self.lock = threading.Lock()

    def allow(self) -> bool:
        with self.lock:
            if self.tokens > 0:
                self.tokens -= 1
                return True
            return False
```

> **Python differs**: use `threading.Lock()` for thread safety (same concept as Go's sync.Mutex).

## Tiny Example

- `main.go` -- token bucket rate limiter middleware, returns 429 when exhausted
- `main.py` -- same with threading.Lock-based limiter

## Common Interview Traps

- **Not thread-safe**: rate limiter shared across goroutines must use mutex
- **No refill mechanism**: tokens must replenish over time or the limiter permanently blocks
- **Global vs per-IP**: global limits all clients together; per-IP is fairer but needs a map
- **No 429 status code**: must return HTTP 429 Too Many Requests, not 403 or 500
- **No Retry-After header**: good practice to tell clients when to retry

## What to Say in Interviews

- "I implement a token bucket with a mutex -- Allow() decrements tokens, returns false when empty"
- "Tokens refill on a timer so the limiter resets over time"
- "I return 429 with a Retry-After header when the limit is exceeded"

## Run It

```bash
go run ./08_http_and_backend/09_rate_limit_basics/
# Rapidly hit: curl http://127.0.0.1:PORT/api  (first 5 succeed, then 429)

python ./08_http_and_backend/09_rate_limit_basics/main.py
```

## TL;DR (Interview Summary)

- Token bucket: fixed number of tokens, consumed per request, refilled over time
- Must be concurrency-safe: `sync.Mutex` in Go, `threading.Lock` in Python
- Return HTTP 429 Too Many Requests when limit exceeded
- Include `Retry-After` header with seconds until refill
- Global rate limiter: one bucket for all clients
- Per-IP rate limiter: map of buckets keyed by client IP (needs cleanup)
- Common algorithms: token bucket, fixed window, sliding window log
