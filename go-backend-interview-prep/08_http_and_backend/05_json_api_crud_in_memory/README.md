# JSON API CRUD (In-Memory)

## What It Is

- **CRUD**: Create, Read, Update, Delete -- the four basic operations on a resource
- A JSON API exposes CRUD via HTTP methods: POST, GET, PUT/PATCH, DELETE

## Why It Matters

- Building a CRUD API from scratch is the most common backend interview exercise
- Interviewers check struct design, mutex safety, status codes, and error handling

## Syntax Cheat Sheet

```go
// Go: typical in-memory CRUD pattern
var (
    store = map[string]Item{}
    mu    sync.RWMutex
)

// POST /items       -> create
// GET  /items       -> list all
// GET  /items/{id}  -> get one
// PUT  /items/{id}  -> update
// DELETE /items/{id} -> delete
```

```python
# Python: dict-based store
items: dict[str, dict] = {}

# Same HTTP method -> operation mapping
```

> **Python differs**: no sync.RWMutex needed for single-threaded servers.
> For threaded servers, use `threading.Lock()`.

## Tiny Example

- `main.go` -- full CRUD API for "items" with mutex-protected map, proper status codes
- `main.py` -- same with dict store and BaseHTTPRequestHandler

## Common Interview Traps

- **No mutex on shared state**: concurrent requests corrupt the map
- **Wrong status codes**: 201 for create, 200 for read/update, 204 or 200 for delete, 404 for missing
- **Forgetting Content-Type header**: always set `application/json`
- **Not closing request body**: Go's `json.NewDecoder` reads the body, but if you skip decoding, close it
- **ID generation**: use a counter or UUID -- don't let clients pick IDs for create

## What to Say in Interviews

- "I protect the store with sync.RWMutex -- RLock for reads, Lock for writes"
- "I return 201 Created with the new resource, 404 for missing items, 400 for bad input"
- "I keep the handler thin -- parse, validate, call store, return JSON"

## Run It

```bash
go run ./08_http_and_backend/05_json_api_crud_in_memory/
# curl http://127.0.0.1:PORT/items
# curl -X POST -H 'Content-Type: application/json' \
#   -d '{"name":"widget","price":9.99}' http://127.0.0.1:PORT/items
# curl http://127.0.0.1:PORT/items/1
# curl -X PUT -H 'Content-Type: application/json' \
#   -d '{"name":"gadget","price":19.99}' http://127.0.0.1:PORT/items/1
# curl -X DELETE http://127.0.0.1:PORT/items/1

python ./08_http_and_backend/05_json_api_crud_in_memory/main.py
```

## TL;DR (Interview Summary)

- CRUD maps to POST (create), GET (read), PUT (update), DELETE (delete)
- Protect shared state with `sync.RWMutex` -- RLock for reads, Lock for writes
- Status codes: 201 create, 200 read/update, 200 delete, 404 not found, 400 bad input
- Always set `Content-Type: application/json`
- Server generates IDs -- clients don't pick them
- Keep handlers thin: parse -> validate -> store operation -> respond
- Python: `threading.Lock()` for thread safety; `dict` as store
