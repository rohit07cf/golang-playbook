# Graceful Shutdown

## What It Is

- **Graceful shutdown**: stop accepting new connections, finish in-flight requests, then exit
- Go: `http.Server.Shutdown(ctx)` does this; `server.Close()` is abrupt (drops connections)

## Why It Matters

- Abrupt shutdown drops in-flight requests -- users see errors, data may be lost
- Interviewers expect you to demonstrate signal handling + graceful drain

## Syntax Cheat Sheet

```go
// Go: signal + graceful shutdown
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

go func() {
    <-quit
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    server.Shutdown(ctx)  // drains in-flight, then stops
}()

server.ListenAndServe()  // blocks until Shutdown is called
```

```python
# Python: signal handling + server shutdown
import signal

def handler(signum, frame):
    server.shutdown()  # stops serve_forever loop

signal.signal(signal.SIGINT, handler)
server.serve_forever()
```

> **Python differs**: `HTTPServer.shutdown()` stops `serve_forever()` but doesn't drain
> in-flight requests as gracefully as Go's `Shutdown(ctx)`.

## Tiny Example

- `main.go` -- server with signal handling, graceful shutdown with timeout
- `main.py` -- similar with signal.signal() handler

## Common Interview Traps

- **Using server.Close() instead of Shutdown()**: Close is immediate; Shutdown drains
- **No shutdown timeout**: if a handler hangs, Shutdown blocks forever -- use context.WithTimeout
- **Forgetting signal.Notify**: without it, SIGINT/SIGTERM kill the process instantly
- **Buffered channel for signal**: `make(chan os.Signal, 1)` -- unbuffered may miss the signal

## What to Say in Interviews

- "I catch SIGINT/SIGTERM, then call server.Shutdown(ctx) with a timeout to drain in-flight requests"
- "Shutdown stops accepting new connections and waits for active ones to complete"
- "I use a context timeout as a safety net -- if requests don't finish in 10 seconds, force stop"

## Run It

```bash
go run ./08_http_and_backend/10_graceful_shutdown/
# Send SIGINT (Ctrl+C) to see graceful shutdown
# Or just wait for demo auto-shutdown

python ./08_http_and_backend/10_graceful_shutdown/main.py
```

## TL;DR (Interview Summary)

- `server.Shutdown(ctx)` -- graceful: stops new connections, drains in-flight
- `server.Close()` -- abrupt: drops everything
- Catch `SIGINT`/`SIGTERM` with `signal.Notify(ch, ...)` on a buffered channel
- Use `context.WithTimeout` for shutdown deadline -- don't hang forever
- Shutdown blocks until all connections close or context expires
- Python: `server.shutdown()` stops the loop but doesn't drain as gracefully
- Always log when shutdown starts and completes for observability
