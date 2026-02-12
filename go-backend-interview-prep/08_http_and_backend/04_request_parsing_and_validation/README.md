# Request Parsing and Validation

## What It Is

- **Request parsing**: decoding JSON body, query params, and headers from an incoming HTTP request
- **ELI10:** Parsing a request is like opening a letter and checking the sender wrote their address correctly.
- **Validation**: checking required fields, types, and constraints before processing

## Why It Matters

- Every API must reject malformed input with clear 400 errors
- **ELI10:** Trusting user input without validation is like eating mystery food from a stranger -- always check before you swallow.
- Interviewers test whether you validate early and return structured error responses

## Syntax Cheat Sheet

```go
// Go: decode JSON body + validate
var req struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, `{"error":"invalid json"}`, 400)
    return
}
if req.Name == "" {
    http.Error(w, `{"error":"name required"}`, 400)
    return
}

// Query params
q := r.URL.Query().Get("page")
```

```python
# Python: parse JSON body
length = int(self.headers.get("Content-Length", 0))
body = json.loads(self.rfile.read(length))
name = body.get("name", "")
if not name:
    # return 400 error
```

> **Python differs**: no struct tags for decoding; you read and validate dicts manually.

## Tiny Example

- `main.go` -- POST endpoint that parses JSON, validates fields, returns 400 on bad input
- `main.py` -- same with manual dict validation

## Common Interview Traps

- **No Content-Type check**: accepting non-JSON bodies silently fails
- **Missing body limit**: `json.NewDecoder(r.Body)` reads unlimited data -- use `http.MaxBytesReader`
- **Ignoring decode errors**: always check the error from Decode
- **Empty vs missing fields**: `""` is valid but empty -- validate both presence and content
- **Query param types**: `r.URL.Query().Get()` returns string -- convert and validate manually

## What to Say in Interviews

- "I decode the JSON body with json.NewDecoder, check the error, then validate each field"
- "I use http.MaxBytesReader to limit body size and prevent denial-of-service"
- "I return structured JSON errors with specific field names so clients know what to fix"

## Run It

```bash
go run ./08_http_and_backend/04_request_parsing_and_validation/
# curl -X POST -H 'Content-Type: application/json' \
#   -d '{"name":"Alice","email":"a@b.com"}' http://127.0.0.1:PORT/users
# curl -X POST -d '{}' http://127.0.0.1:PORT/users   # -> 400

python ./08_http_and_backend/04_request_parsing_and_validation/main.py
```

## TL;DR (Interview Summary)

- `json.NewDecoder(r.Body).Decode(&v)` -- parse JSON into struct
- Always check decode error -- return 400 with message
- Validate required fields, lengths, formats before processing
- `http.MaxBytesReader(w, r.Body, limit)` -- prevent huge payloads
- `r.URL.Query().Get("key")` -- query params (always strings)
- Return structured JSON errors: `{"error":"name required"}`
- Python: `json.loads(self.rfile.read(length))` + manual validation
