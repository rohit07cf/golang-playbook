# 05 -- Observability Basics (Mini)

## What We Are Building

- An HTTP service with request IDs, structured logging, latency tracking, and a `/metrics` endpoint
- **ELI10:** Observability is putting dashcams in every car in your fleet -- when something crashes, you rewind the tape.
- A toy observability layer using only the standard library

## Requirements

**Functional:**
- Every request gets a unique request ID (UUID-ish)
- Structured log lines: timestamp, request_id, method, path, status, latency
- In-memory metrics: request count, error count, latency histogram (simple)
- `GET /metrics` exposes counters as plain text

**Non-functional:**
- Minimal overhead per request
- Concurrency-safe metric counters
- Request ID propagated via context

## High-Level Design

```
Client --> [RequestID middleware] --> [Logging middleware] --> Handler
                                          |
                                    writes to Metrics
                                          |
                                    GET /metrics --> dump counters
```

```
+--------+     +------------+     +----------+     +---------+
| Client | --> | ReqID MW   | --> | Log MW   | --> | Handler |
+--------+     +------------+     +----------+     +---------+
                                       |
                                  +---------+
                                  | Metrics |
                                  | (atomic)|
                                  +---------+
                                       |
                               GET /metrics
```

## Key Go Building Blocks Used

- `context.WithValue` -- propagate request ID through the call chain
- `sync/atomic` -- lock-free metric counters
- `time.Since` -- measure request latency
- `crypto/rand` -- generate request IDs
- Middleware chaining -- compose handlers

## Trade-Offs

- **In-memory metrics** -- lost on restart; production uses Prometheus/Datadog
- **Atomic counters only** -- no histograms or percentiles; production uses proper libraries
- **Request ID in context** -- clean but adds allocations; acceptable overhead
- **Plain text /metrics** -- toy format; production uses Prometheus exposition format
- **No distributed tracing** -- single service only; production uses OpenTelemetry

## Common Interview Traps

- Not propagating request ID across service boundaries
- Logging too much (every field) or too little (no context)
- Using mutex for counters when atomic is sufficient
- Forgetting latency percentiles (p50, p95, p99) -- averages hide outliers
- Not mentioning the three pillars: logs, metrics, traces

## Run It

```bash
go run ./11_system_design_in_go/05_observability_basics_mini
python3 ./11_system_design_in_go/05_observability_basics_mini/main.py
```

## TL;DR

- Every request gets a unique ID for tracing through logs
- Structured logging: timestamp, request_id, method, path, status, latency_ms
- `sync/atomic` for lock-free metric counters
- `context.WithValue` propagates request ID without polluting function signatures
- `/metrics` endpoint exposes counters (request count, error count, latency)
- Three pillars of observability: **logs**, **metrics**, **traces**
- In interviews, always mention you would use Prometheus + OpenTelemetry in production
