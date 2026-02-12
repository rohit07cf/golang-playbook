# Logging and Request ID

## What It Is

- **Request ID**: a unique identifier (UUID or random hex) attached to every request for tracing
- **Structured logging**: log lines that include request ID, method, path, status, duration

## Why It Matters

- Without request IDs, correlating logs across services is impossible
- Interviewers expect you to show middleware that injects a request ID into context and logs

## Syntax Cheat Sheet

```go
// Go: generate request ID, store in context
func requestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := uuid.New().String()  // or crypto/rand hex
        ctx := context.WithValue(r.Context(), "requestID", id)
        w.Header().Set("X-Request-ID", id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

```python
# Python: generate request ID in handler
import uuid
request_id = str(uuid.uuid4())
```

> **Python differs**: no context propagation -- pass request ID explicitly or use thread-local storage.

## Tiny Example

- `main.go` -- request ID middleware + logging middleware that prints structured log lines
- `main.py` -- same with manual request ID generation per request

## Common Interview Traps

- **Using context.WithValue with string key**: use a private type to avoid collisions
- **Not returning the ID to the client**: set `X-Request-ID` header so clients can reference it
- **Logging after response**: log at the end of the middleware to capture duration and status
- **No structured format**: use key=value pairs or JSON -- not free-form strings

## What to Say in Interviews

- "I inject a request ID via middleware and store it in context for downstream use"
- "Every log line includes the request ID so I can trace a request across services"
- "I use a typed context key to avoid string collisions in context.WithValue"

## Run It

```bash
go run ./08_http_and_backend/07_logging_and_request_id/
# curl -v http://127.0.0.1:PORT/hello   # check X-Request-ID header

python ./08_http_and_backend/07_logging_and_request_id/main.py
```

## TL;DR (Interview Summary)

- Generate unique ID per request (crypto/rand or UUID)
- Store in context: `context.WithValue(ctx, key, id)` -- use typed key, not string
- Set `X-Request-ID` response header for client tracing
- Log: method, path, status, duration, request ID -- structured format
- Middleware pattern: generate ID -> inject into context -> call next -> log on return
- Python: no context propagation -- thread-local or explicit parameter passing
- Accept incoming `X-Request-ID` header to support distributed tracing
