# Auth Basics (API Key)

## What It Is

- **API key auth**: client sends a secret key in a header; server validates it before processing
- **ELI10:** API key auth is a bouncer checking your wristband -- no band, no entry, no exceptions.
- Simplest auth scheme -- good for internal services; not suitable for user-facing auth alone

## Why It Matters

- Interviewers test whether you can write auth middleware that rejects unauthenticated requests
- **ELI10:** Without auth, your API is a house with no locks -- anyone can walk in and rearrange the furniture.
- Shows understanding of the handler chain: check auth first, then proceed

## Syntax Cheat Sheet

```go
// Go: API key middleware
func apiKeyAuth(key string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-API-Key") != key {
            http.Error(w, `{"error":"unauthorized"}`, 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

```python
# Python: check header in handler
api_key = self.headers.get("X-API-Key", "")
if api_key != EXPECTED_KEY:
    json_response(self, 401, {"error": "unauthorized"})
    return
```

> **Note**: This is a toy example. Production auth uses JWT, OAuth, or session tokens.
> API keys are fine for service-to-service calls but not for user authentication.

## Tiny Example

- `main.go` -- API key middleware protecting endpoints; returns 401 JSON on failure
- `main.py` -- same with header check in handler

## Common Interview Traps

- **Hardcoded keys in source**: use environment variables or config files
- **Timing attacks**: use `subtle.ConstantTimeCompare` for key comparison
- **No JSON error body**: return structured `{"error":"unauthorized"}`, not plain text
- **Auth on wrong endpoints**: public endpoints (health check) shouldn't require auth
- **Confusing auth with authz**: authentication = who you are; authorization = what you can do

## What to Say in Interviews

- "I implement API key auth as middleware that checks the X-API-Key header before calling next"
- "I use constant-time comparison to prevent timing attacks on the key"
- "This is a simple scheme for service-to-service auth; for user auth I'd use JWT or OAuth"

## Run It

```bash
go run ./08_http_and_backend/08_auth_basics_api_key/
# curl http://127.0.0.1:PORT/protected             # -> 401
# curl -H 'X-API-Key: secret123' http://127.0.0.1:PORT/protected  # -> 200

python ./08_http_and_backend/08_auth_basics_api_key/main.py
```

## TL;DR (Interview Summary)

- API key: simple header-based auth (`X-API-Key`)
- Middleware checks key before calling next handler
- `crypto/subtle.ConstantTimeCompare` -- prevent timing attacks
- Return 401 JSON `{"error":"unauthorized"}` -- not empty body
- Don't hardcode keys -- use env vars or config
- Public endpoints (health) skip auth middleware
- This is a toy -- production uses JWT/OAuth for user auth
