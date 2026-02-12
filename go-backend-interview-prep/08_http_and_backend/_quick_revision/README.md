# HTTP & Backend -- Quick Revision

> One-screen refresher. Skim before your interview.

## 9 Tiny Go Snippets

```go
// 1. Basic HTTP server
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
})
http.ListenAndServe(":8080", mux)
```

```go
// 2. Routing with path params (Go 1.22+)
mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    json.NewEncoder(w).Encode(map[string]string{"id": id})
})
```

```go
// 3. Middleware pattern
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
// Use: http.ListenAndServe(":8080", logging(mux))
```

```go
// 4. Parse + validate JSON body
var req struct{ Name string `json:"name"` }
r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "bad json", 400); return
}
```

```go
// 5. Context timeout per request
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()
select {
case res := <-doWork(ctx): json.NewEncoder(w).Encode(res)
case <-ctx.Done(): http.Error(w, "timeout", 504)
}
```

```go
// 6. Request ID middleware
func reqID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := hex.EncodeToString(randomBytes(8))
        ctx := context.WithValue(r.Context(), "rid", id)
        w.Header().Set("X-Request-ID", id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

```go
// 7. API key auth middleware
func auth(key string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if subtle.ConstantTimeCompare([]byte(r.Header.Get("X-API-Key")), []byte(key)) != 1 {
            http.Error(w, `{"error":"unauthorized"}`, 401); return
        }
        next.ServeHTTP(w, r)
    })
}
```

```go
// 8. Graceful shutdown
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
go func() { <-quit; server.Shutdown(context.Background()) }()
server.ListenAndServe()
```

```go
// 9. Repository interface
type Repo interface {
    Create(name string) (Item, error)
    GetByID(id int) (Item, error)
}
// Swap: MemoryRepo{} or SQLiteRepo{} -- handlers don't change
```

## 9 Tiny Python Snippets

```python
# 1. Basic HTTP server
from http.server import HTTPServer, BaseHTTPRequestHandler
class H(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200); self.end_headers()
        self.wfile.write(b"hello")
HTTPServer(("", 8080), H).serve_forever()
```

```python
# 2. Manual routing (no framework)
def do_GET(self):
    if self.path == "/items":
        ...  # list
    elif self.path.startswith("/items/"):
        item_id = self.path.split("/")[-1]
```

```python
# 3. Middleware (wrap in try/except)
def do_GET(self):
    start = time.perf_counter()
    try: self._route()
    finally: print(f"{self.command} {self.path} {time.perf_counter()-start:.3f}s")
```

```python
# 4. Parse + validate JSON
length = int(self.headers.get("Content-Length", 0))
body = json.loads(self.rfile.read(length))
if not body.get("name"):
    json_response(self, 400, {"error": "name required"}); return
```

```python
# 5. Timeout via threading
cancel = threading.Event()
t = threading.Thread(target=worker, args=(cancel,))
t.start(); t.join(timeout=2.0)
if t.is_alive(): cancel.set()  # timed out
```

```python
# 6. Request ID
import os
request_id = os.urandom(8).hex()
self.send_header("X-Request-ID", request_id)
```

```python
# 7. API key check
import hmac
if not hmac.compare_digest(self.headers.get("X-API-Key",""), KEY):
    json_response(self, 401, {"error": "unauthorized"}); return
```

```python
# 8. Graceful shutdown
import signal
signal.signal(signal.SIGINT, lambda s,f: threading.Thread(target=server.shutdown).start())
server.serve_forever()
```

```python
# 9. SQLite repository (stdlib)
import sqlite3
conn = sqlite3.connect(":memory:")
conn.execute("CREATE TABLE items (id INTEGER PRIMARY KEY, name TEXT)")
conn.execute("INSERT INTO items (name) VALUES (?)", ("widget",))
```

## 9 Interview One-Liners

| # | Topic | One-Liner |
|---|-------|-----------|
| 1 | HTTP server | `http.ListenAndServe(":8080", mux)` -- handler func takes `(w, r)` |
| 2 | Routing | Go 1.22+: `"GET /items/{id}"` -- `r.PathValue("id")` for params |
| 3 | Middleware | `func(next http.Handler) http.Handler` -- chain: `logging(auth(mux))` |
| 4 | Parsing | `json.NewDecoder(r.Body).Decode(&v)` + `MaxBytesReader` for safety |
| 5 | Context | `context.WithTimeout` per request -- `defer cancel()` always |
| 6 | Request ID | Generate, store in context, set `X-Request-ID` header |
| 7 | Auth | `subtle.ConstantTimeCompare` for API key; return 401 JSON |
| 8 | Shutdown | `server.Shutdown(ctx)` drains in-flight; `Close()` is abrupt |
| 9 | Repo pattern | Interface abstracts storage -- swap impl without changing handlers |

## TL;DR

- **Server**: `http.NewServeMux` + `HandleFunc` + `ListenAndServe`
- **Routing**: Go 1.22+ method+path patterns; Python: manual `self.path` parsing
- **Middleware**: `func(http.Handler) http.Handler` -- composable chain
- **Parsing**: Decode JSON, validate, return 400 with structured errors
- **Context**: `WithTimeout` + `defer cancel()` -- bound every request
- **Request ID**: middleware generates, injects into context + response header
- **Auth**: API key in header, constant-time compare, 401 on failure
- **Rate limit**: token bucket with mutex, return 429 + Retry-After
- **Shutdown**: `Shutdown(ctx)` drains; catch SIGINT/SIGTERM
- **Repo**: interface -> implementation -> service -> handler layering
