# System Design in Go -- Quick Revision

One-screen cheat sheet for system design interviews.

---

## The 7-Step Answer Framework

1. **Clarify** -- ask scope, users, scale
2. **Functional requirements** -- what the system does (APIs)
3. **Non-functional requirements** -- latency, throughput, availability
4. **Back-of-envelope** -- requests/sec, storage, bandwidth
5. **High-level design** -- boxes: client, LB, services, DB, cache, queue
6. **Deep-dive** -- pick the hardest component, design it thoroughly
7. **Trade-offs** -- what breaks at 10x, what you would change

---

## Patterns to Mention

| Pattern | When to Use | Go Hook |
|---------|------------|---------|
| **Rate limiting** | Protect services from abuse | Token bucket + `sync.Mutex` |
| **Caching** | Reduce DB load, speed up reads | Map + TTL + LRU eviction |
| **Job queue** | Async processing, decouple producer/consumer | Buffered `chan` + worker goroutines |
| **Retry + backoff** | Handle transient failures | `time.Sleep` + exponential + jitter |
| **Circuit breaker** | Fail fast when downstream is dead | State machine + `sync.Mutex` |
| **Timeout** | Prevent hanging on slow calls | `context.WithTimeout` |
| **Observability** | Debug and monitor in production | Request ID + structured logs + `/metrics` |

---

## Go Implementation Hooks

- **Middleware**: `func(http.Handler) http.Handler` -- compose auth, rate limit, logging
- **Interfaces**: define at consumer, implement at provider -- swap DB, cache, queue
- **context.Context**: propagate timeouts, cancellation, request ID across call chain
- **Worker pools**: buffered channel + N goroutines + `sync.WaitGroup`
- **Atomic counters**: `sync/atomic` for lock-free metrics
- **sync.RWMutex**: concurrent reads, exclusive writes (cache, store)
- **Channel close**: signal "no more work" to consumer goroutines

---

## 7 Interview One-Liners

| # | One-Liner |
|---|-----------|
| 1 | "I would use a **token bucket** for rate limiting because it handles bursts gracefully." |
| 2 | "For the URL shortener, I would use **base62 encoding** with collision retry." |
| 3 | "The job queue uses a **buffered channel** for backpressure -- producer blocks when full." |
| 4 | "I would add a **TTL-based cache** in front of the DB to reduce read latency." |
| 5 | "Every request gets a **unique ID** propagated via context for distributed tracing." |
| 6 | "I would wrap downstream calls with **retry + circuit breaker** -- retry for transient failures, circuit breaker to fail fast when the service is down." |
| 7 | "The API gateway is the single entry point -- it handles **auth, rate limiting, routing, and observability** as cross-cutting middleware." |

---

## TL;DR

- Always follow the 7-step framework -- interviewers notice structure
- Mention caching, rate limiting, and queues early -- they solve most scaling problems
- Go stdlib covers all building blocks: `net/http`, `context`, `sync`, `chan`
- Interfaces + middleware = clean, testable, composable architecture
- Call out trade-offs explicitly -- that is what separates senior from junior answers
- Practice explaining each pattern in 30 seconds with a concrete example
