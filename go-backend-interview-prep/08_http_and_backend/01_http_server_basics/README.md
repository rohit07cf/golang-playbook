# HTTP Server Basics

## What It Is

- `http.ListenAndServe(addr, handler)` starts a blocking HTTP server
- A **handler** is any function matching `func(w http.ResponseWriter, r *http.Request)`

## Why It Matters

- Every Go backend starts here -- no framework needed
- Interviewers expect you to build an HTTP server from scratch with `net/http`

## Syntax Cheat Sheet

```go
// Go: minimal server
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
})
http.ListenAndServe(":8080", nil)
```

```python
# Python: http.server
from http.server import HTTPServer, BaseHTTPRequestHandler
class H(BaseHTTPRequestHandler):
    def do_GET(self): self.send_response(200); self.end_headers(); self.wfile.write(b"hello")
HTTPServer(("", 8080), H).serve_forever()
```

> **Python differs**: `http.server` requires subclassing `BaseHTTPRequestHandler`.
> Go uses plain functions. Both are stdlib-only.

## Tiny Example

- `main.go` -- `/hello` and `/health` endpoints, auto-shutdown after a few seconds for demo
- `main.py` -- same endpoints using `http.server`

## Common Interview Traps

- **`nil` mux uses DefaultServeMux**: fine for demos, avoid in production (global state)
- **ListenAndServe blocks**: it never returns unless there's an error
- **ResponseWriter is write-once for headers**: calling WriteHeader twice is a bug
- **Forgetting Content-Type**: Go sets `text/plain` by default; set `application/json` explicitly

## What to Say in Interviews

- "I use net/http with explicit ServeMux -- no global DefaultServeMux in production"
- "A handler is just `func(ResponseWriter, *Request)` -- the simplest interface possible"
- "For production I create an `http.Server` struct for timeout and shutdown control"

## Run It

```bash
go run ./08_http_and_backend/01_http_server_basics/
# In another terminal:
curl http://localhost:8080/hello
curl http://localhost:8080/health
```

```bash
python ./08_http_and_backend/01_http_server_basics/main.py
# In another terminal:
curl http://localhost:8081/hello
curl http://localhost:8081/health
```

## TL;DR (Interview Summary)

- `http.HandleFunc(pattern, handlerFunc)` registers a route
- `http.ListenAndServe(":8080", nil)` starts the server (blocks)
- Handler signature: `func(w http.ResponseWriter, r *http.Request)`
- `w.Header().Set(...)` before `w.Write(...)` or `w.WriteHeader(...)`
- Use `http.Server{}` struct in production for timeouts + shutdown
- Python: subclass `BaseHTTPRequestHandler`, override `do_GET` / `do_POST`
