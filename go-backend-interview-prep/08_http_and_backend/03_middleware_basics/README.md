# Middleware Basics

## What It Is

- **Middleware**: a function that wraps an HTTP handler, runs code before/after the actual handler
- **ELI10:** Middleware is the security checkpoint at the airport -- every request passes through it before reaching the gate.
- Go pattern: `func(next http.Handler) http.Handler` -- returns a new handler that calls `next`

## Why It Matters

- Every production server uses middleware for logging, auth, CORS, panic recovery
- **ELI10:** Without middleware you'd copy-paste the same logging/auth code into every handler -- middleware lets you write it once and wrap everything.
- Interviewers expect you to explain the chain pattern and write one from scratch

## Syntax Cheat Sheet

```go
// Go: middleware signature
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        next.ServeHTTP(w, r)          // call next handler
    })
}

// Chain: logging(auth(mux))
```

```python
# Python: decorator-style middleware
def logging_mw(handler_class):
    class Wrapped(handler_class):
        def do_GET(self):
            print(f"{self.command} {self.path}")
            super().do_GET()
    return Wrapped
```

> **Python differs**: no standard middleware chain in stdlib.
> You subclass or wrap handler classes manually. Frameworks like WSGI/ASGI have standard middleware.

## Tiny Example

- `main.go` -- logging + timing middleware chained around a simple handler
- `main.py` -- same with decorator-style wrapping

## Common Interview Traps

- **Wrong chain order**: middleware executes outside-in; `logging(auth(mux))` logs first, then checks auth
- **Forgetting to call next**: middleware that doesn't call `next.ServeHTTP` swallows the request
- **Response already written**: if middleware writes headers, the inner handler can't change status code
- **Panic in middleware**: always defer a recover in production middleware

## What to Say in Interviews

- "I write middleware as `func(http.Handler) http.Handler` and chain them: `logging(auth(mux))`"
- "Middleware runs outside-in -- the outermost runs first on the request and last on the response"
- "I keep middleware small and composable -- each one does exactly one thing"

## Run It

```bash
go run ./08_http_and_backend/03_middleware_basics/
# In another terminal:
#   curl http://127.0.0.1:PORT/hello

python ./08_http_and_backend/03_middleware_basics/main.py
```

## TL;DR (Interview Summary)

- Middleware signature: `func(next http.Handler) http.Handler`
- Chain: `logging(auth(handler))` -- outside-in execution
- Always call `next.ServeHTTP(w, r)` -- or the chain breaks
- Use for: logging, auth, CORS, rate limiting, panic recovery
- Keep each middleware single-purpose and composable
- Python stdlib has no middleware chain -- subclass or use WSGI
- Execution order: request flows in, response flows back out through the chain
