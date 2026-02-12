# 01 -- Rate Limiter Service (Mini)

## What We Are Building

- An HTTP service with a `GET /ping` endpoint
- Middleware enforces **per-client rate limiting** using a token bucket algorithm

## Requirements

**Functional:**
- `GET /ping` returns `{"message": "pong"}` if within rate limit
- Return `429 Too Many Requests` when limit is exceeded
- Rate limit is per-client (keyed by IP or header)

**Non-functional:**
- Concurrency-safe (multiple goroutines hitting the same limiter)
- Low latency -- the middleware should add negligible overhead
- Configurable rate and burst size

## High-Level Design

```
Client --> [RateLimitMiddleware] --> /ping handler --> JSON response
               |
         per-client bucket
         (map[string]*Bucket)
         protected by sync.Mutex
```

- Each client IP gets its own token bucket
- Bucket refills at a fixed rate (e.g., 2 tokens/sec)
- Burst allows short spikes (e.g., 5 tokens max)
- Middleware checks bucket before forwarding to handler

## Key Go Building Blocks Used

- `net/http` -- HTTP server + handler chaining
- `sync.Mutex` -- protect shared bucket map
- `time.Now()` -- calculate token refill
- `http.HandlerFunc` wrapping -- middleware pattern
- `encoding/json` -- JSON responses

## Trade-Offs

- **In-memory only** -- does not survive restarts, not shared across instances
- **Token bucket vs fixed window** -- token bucket allows bursts, fixed window is simpler
- **Per-IP keying** -- fails behind shared proxies (use API key header instead)
- **Single mutex** -- works for moderate load; at high scale, use sharded maps or sync.Map
- **No distributed coordination** -- for multi-instance, use Redis-based rate limiting

## Common Interview Traps

- Forgetting to make the bucket map concurrency-safe
- Using a single global counter instead of per-client tracking
- Not explaining the difference between token bucket, leaky bucket, and fixed window
- Ignoring clock drift in distributed rate limiting
- Not mentioning what happens when rate limit state is lost (restart tolerance)

## Run It

```bash
go run ./11_system_design_in_go/01_rate_limiter_service_mini
python3 ./11_system_design_in_go/01_rate_limiter_service_mini/main.py
```

## TL;DR

- Token bucket: refill tokens at fixed rate, each request costs one token
- Per-client keying: use IP, API key, or user ID as the bucket key
- `sync.Mutex` protects the shared map of buckets
- Middleware pattern: wrap the real handler, check limit before forwarding
- Return 429 with `Retry-After` header when exhausted
- In production: use Redis for shared state across instances
- Always mention the algorithm name in interviews ("I would use a token bucket because...")
