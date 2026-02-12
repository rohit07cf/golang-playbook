# 08 -- HTTP and Backend

This module builds a complete backend skill set using **only** the
Go and Python standard libraries. No frameworks -- just `net/http`,
`encoding/json`, `context`, and their Python equivalents.

You'll build handlers, middleware chains, JSON APIs, authentication,
rate limiting, graceful shutdown, and a repository-pattern service.
These are the exact patterns interviewers expect in system design rounds.

Each example includes `main.go` + Python equivalent `main.py`.

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `01_http_server_basics` | http.ListenAndServe, handler func, health endpoint |
| 02 | `02_routing_and_handlers` | ServeMux routing, path parameters, method checks |
| 03 | `03_middleware_basics` | Middleware chain, logging middleware, timing middleware |
| 04 | `04_request_parsing_and_validation` | JSON body parsing, field validation, 400 errors |
| 05 | `05_json_api_crud_in_memory` | Full CRUD API, in-memory store with mutex |
| 06 | `06_context_timeouts_cancel` | Per-request timeouts, context cancellation |
| 07 | `07_logging_and_request_id` | Request ID generation, propagation, structured logging |
| 08 | `08_auth_basics_api_key` | API key middleware, 401 responses |
| 09 | `09_rate_limit_basics` | Token bucket rate limiter, 429 responses |
| 10 | `10_graceful_shutdown` | http.Server.Shutdown, signal handling |
| 11 | `11_sqlite_repo_pattern_intro` | Repository pattern, service layer, handler layer |

---

## 10-Min Revision Path

1. Skim `01_http_server_basics` -- handler signature, ListenAndServe
2. Skim `02_routing_and_handlers` -- ServeMux, method routing
3. Skim `03_middleware_basics` -- the chain pattern
4. Skim `05_json_api_crud_in_memory` -- the CRUD API interviewers love
5. Skim `06_context_timeouts_cancel` -- per-request timeout pattern
6. Skim `08_auth_basics_api_key` -- auth middleware shape
7. Skim `10_graceful_shutdown` -- Shutdown(ctx) pattern
8. Skim `_quick_revision/` -- one-screen cheat sheet

---

## Common HTTP / Backend Mistakes

- Using the default `http.Client` (no timeout -- hangs forever)
- Forgetting to set `Content-Type: application/json` on responses
- Not closing `resp.Body` on the client side (leaks connections)
- Reading the entire request body without a size limit (OOM attack)
- Not protecting shared state with a mutex (data race in CRUD handlers)
- Forgetting `defer cancel()` after `context.WithTimeout`
- No graceful shutdown -- in-flight requests get killed
- Hardcoding secrets in source (use env vars)
- Not validating input at the handler boundary
- Returning 200 for errors instead of proper status codes

---

## TL;DR

- `http.HandleFunc("/path", handler)` + `http.ListenAndServe(":8080", nil)`
- Middleware = `func(http.Handler) http.Handler` -- chain them
- Always parse + validate JSON at the boundary, return 400 with details
- Protect shared state with `sync.Mutex` in handlers
- Set per-request timeouts with `context.WithTimeout`
- Generate request IDs in middleware, propagate via context
- Graceful shutdown: `server.Shutdown(ctx)` on SIGINT/SIGTERM
- Repository pattern: interface -> in-memory/DB impl -> service -> handler
