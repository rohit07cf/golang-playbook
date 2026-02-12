# Routing and Handlers

## What It Is

- **Routing**: mapping URL paths + HTTP methods to handler functions
- **ELI10:** A router is the receptionist -- "You want /users? Door 2. You want /orders? Door 5."
- Go 1.22+ `ServeMux` supports `GET /items/{id}` path patterns; older versions need manual checks

## Why It Matters

- Every backend interview expects you to wire up routes correctly
- **ELI10:** Without routing, every request goes to the same place -- like a building with one room and no signs.
- Interviewers check whether you handle method restrictions and 404s cleanly

## Syntax Cheat Sheet

```go
// Go 1.22+: method + path pattern
mux := http.NewServeMux()
mux.HandleFunc("GET /items/{id}", getItem)   // path param
mux.HandleFunc("POST /items", createItem)

// Access path param
id := r.PathValue("id")
```

```python
# Python: manual routing in BaseHTTPRequestHandler
class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path.startswith("/items/"):
            item_id = self.path.split("/")[-1]
            # handle get
```

> **Python differs**: stdlib has no built-in router; you parse `self.path` manually.
> Frameworks like Flask add decorators, but we stick to stdlib here.

## Tiny Example

- `main.go` -- ServeMux with path params, method-specific routes, 404/405 handling
- `main.py` -- manual path parsing with BaseHTTPRequestHandler

## Common Interview Traps

- **Forgetting method checks**: `HandleFunc("/items", h)` matches ALL methods -- restrict with method prefix
- **Trailing slash mismatch**: `/items` vs `/items/` are different routes
- **No 404 handler**: default ServeMux returns empty 404 -- add a catch-all for clean JSON errors
- **Path param parsing**: before Go 1.22, you split the path manually (like Python)

## What to Say in Interviews

- "I register routes with method + path on ServeMux -- GET /items/{id} for reads, POST /items for creates"
- "I return proper 404/405 JSON errors instead of letting the framework return empty responses"
- "For older Go versions I parse path segments manually, but 1.22+ has native path params"

## Run It

```bash
go run ./08_http_and_backend/02_routing_and_handlers/
# In another terminal:
#   curl http://127.0.0.1:PORT/items
#   curl http://127.0.0.1:PORT/items/42
#   curl -X DELETE http://127.0.0.1:PORT/items/42

python ./08_http_and_backend/02_routing_and_handlers/main.py
```

## TL;DR (Interview Summary)

- `mux.HandleFunc("GET /path/{param}", handler)` -- Go 1.22+ routing
- `r.PathValue("id")` -- extract path parameters
- Always restrict by HTTP method -- don't accept everything
- Return JSON 404/405 errors -- not empty bodies
- Python stdlib: parse `self.path` manually -- no built-in router
- Trailing slash matters: `/items` != `/items/`
- Prefer explicit route registration over catch-all + switch
