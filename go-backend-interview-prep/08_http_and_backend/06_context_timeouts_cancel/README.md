# Context Timeouts and Cancellation

## What It Is

- **Context**: Go's `context.Context` carries deadlines, cancellation signals, and request-scoped values
- Per-request timeouts prevent slow handlers from holding resources forever

## Why It Matters

- Production servers must bound request duration -- unbounded requests leak goroutines and connections
- Interviewers expect you to demonstrate context.WithTimeout and select on `ctx.Done()`

## Syntax Cheat Sheet

```go
// Go: per-request timeout
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()

select {
case result := <-doWork(ctx):
    json.NewEncoder(w).Encode(result)
case <-ctx.Done():
    http.Error(w, "timeout", http.StatusGatewayTimeout)
}
```

```python
# Python: no native context -- use threading.Event or asyncio timeout
import signal, functools

# Simple approach: set a socket/request timeout
urllib.request.urlopen(url, timeout=2)
```

> **Python differs**: no `context.Context` in stdlib. Use timeouts on individual operations
> or `asyncio.wait_for()` in async code. Go's context propagation has no direct equivalent.

## Tiny Example

- `main.go` -- handler with context timeout; fast + slow endpoints demonstrate cancellation
- `main.py` -- similar with threading.Timer-based timeout

## Common Interview Traps

- **Forgetting defer cancel()**: leaks context resources
- **Ignoring ctx.Done()**: long operations must check context or they won't cancel
- **Server timeout vs handler timeout**: `http.Server.WriteTimeout` is global; context is per-handler
- **Context value abuse**: don't store business logic in context values -- only request-scoped metadata

## What to Say in Interviews

- "I wrap each request with context.WithTimeout to bound handler duration"
- "Long operations select on ctx.Done() to bail out early on cancellation"
- "I always defer cancel() to free context resources even if the handler returns early"

## Run It

```bash
go run ./08_http_and_backend/06_context_timeouts_cancel/
# In another terminal:
#   curl http://127.0.0.1:PORT/fast   # -> instant
#   curl http://127.0.0.1:PORT/slow   # -> 504 timeout

python ./08_http_and_backend/06_context_timeouts_cancel/main.py
```

## TL;DR (Interview Summary)

- `context.WithTimeout(r.Context(), duration)` -- per-request deadline
- Always `defer cancel()` -- prevents resource leaks
- `select { case <-ctx.Done(): }` -- bail out on timeout
- Server-level timeouts: `http.Server{ReadTimeout, WriteTimeout}`
- Don't abuse `context.WithValue` -- only for request-scoped metadata (request ID, auth token)
- Python: no context equivalent -- use operation-level timeouts
- Unbounded handlers leak goroutines -- always set deadlines in production
