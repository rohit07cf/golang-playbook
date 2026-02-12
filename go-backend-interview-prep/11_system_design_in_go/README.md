# 11 -- System Design in Go

Build mini systems that demonstrate **system design reasoning** using Go standard library only.

Each folder is a small, runnable service that covers a real interview topic.
No external frameworks -- just `net/http`, `context`, `sync`, `time`, `encoding/json`.
Focus on **clean layering**, **interfaces for DI**, **concurrency safety**, and **trade-offs**.

Every example includes `main.go` + Python equivalent `main.py`.

---

## Mini Systems

| # | Folder | What It Covers |
|---|--------|---------------|
| 00 | `00_system_design_framework` | The 7-step answer framework (README only) |
| 01 | `01_rate_limiter_service_mini` | Token bucket middleware, per-client limiting |
| 02 | `02_url_shortener_service_mini` | Shorten + redirect, in-memory store, collision handling |
| 03 | `03_job_queue_worker_mini` | Channel-based queue, worker pool, backpressure |
| 04 | `04_cache_service_mini` | In-memory cache with TTL, LRU eviction |
| 05 | `05_observability_basics_mini` | Request ID, structured logging, /metrics endpoint |
| 06 | `06_reliability_patterns_mini` | Retry, timeout, circuit breaker |
| 07 | `07_api_gateway_basics_mini` | Reverse proxy routing, auth, rate limit, timeout |

---

## 10-Min Revision Path

1. Read `00_system_design_framework` -- memorize the 7-step structure
2. Run `01_rate_limiter` -- understand middleware + token bucket
3. Run `03_job_queue_worker` -- channels + worker pool pattern
4. Run `04_cache_service` -- TTL + eviction trade-offs
5. Run `06_reliability_patterns` -- retry + circuit breaker combo
6. Skim `_quick_revision/README.md` -- one-screen cheat sheet
7. Practice explaining trade-offs aloud for each mini system

---

## Common System Design Mistakes

- Jumping into implementation before clarifying requirements
- Forgetting non-functional requirements (latency, throughput, durability)
- Using a single global mutex when per-key locking would scale better
- Ignoring backpressure -- unbounded queues cause OOM
- No timeout on downstream calls -- one slow service blocks everything
- Skipping observability -- you cannot debug what you cannot measure
- Over-engineering the first version -- start simple, mention scaling later
- Confusing consistency models -- know when eventual consistency is acceptable
- Not mentioning trade-offs -- interviewers want to see you weigh options

---

## TL;DR

- System design interviews test **structured thinking**, not perfect code
- Always start with requirements, then high-level design, then deep-dive
- Go stdlib is enough to build rate limiters, queues, caches, gateways
- Use **interfaces** to separate layers (handler / service / repo)
- Use **channels** for producer-consumer, **sync.Mutex** for shared state
- Use **context** for timeouts and cancellation propagation
- Call out trade-offs explicitly -- that is what gets you hired
- Every mini system here runs locally with zero setup
