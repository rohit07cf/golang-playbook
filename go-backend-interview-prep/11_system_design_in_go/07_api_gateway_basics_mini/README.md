# 07 -- API Gateway Basics (Mini)

## What We Are Building

- A toy API gateway that routes requests to downstream services
- **ELI10:** An API gateway is the hotel front desk -- all guests check in here, get their room key, and get directed to the right floor.
- Includes auth check, rate limiting, request ID injection, and timeout enforcement

## Requirements

**Functional:**
- Route `/api/users/*` to user service handler
- Route `/api/orders/*` to order service handler
- API key auth check (reject invalid keys)
- Per-client rate limiting
- Request ID on every request

**Non-functional:**
- Timeout on downstream calls (prevent hanging)
- Composable middleware chain
- Clear rejection messages (401, 429, 504)

## High-Level Design

```
Client --> [Auth] --> [RateLimit] --> [RequestID] --> [Timeout] --> Router
                                                                     |
                                                      /api/users  /api/orders
                                                          |            |
                                                     UserHandler  OrderHandler
```

```
+--------+     +------+     +-------+     +-------+     +--------+
| Client | --> | Auth | --> | Rate  | --> | ReqID | --> | Router |
+--------+     | MW   |     | Limit |     | MW    |     |        |
               +------+     | MW    |     +-------+     +--------+
                             +-------+        |          /  \
                                              |         /    \
                                        timeout    Users   Orders
                                        context    Handler  Handler
```

## Key Go Building Blocks Used

- `net/http` -- HTTP server + handler interface
- Middleware chaining -- `func(http.Handler) http.Handler`
- `context.WithTimeout` -- deadline on downstream calls
- `sync.Mutex` -- rate limiter state
- `strings.HasPrefix` -- simple path-based routing

## Trade-Offs

- **Path-based routing** -- simple but no regex; production uses trie-based routers
- **In-process rate limiting** -- not shared; production uses Redis or API gateway service
- **Static API keys** -- toy auth; production uses JWT, OAuth2
- **No load balancing** -- single instance; production uses round-robin, least-connections
- **No request body forwarding** -- simplified; real proxy forwards full request
- **No TLS termination** -- production gateway handles HTTPS

## Common Interview Traps

- Forgetting that the gateway is a single point of failure (needs HA)
- Not mentioning rate limiting at the gateway level (vs per-service)
- Confusing API gateway with load balancer (gateway does routing + auth + transformation)
- Not discussing how timeouts propagate (gateway timeout vs service timeout)
- Ignoring observability at the gateway (it is the best place to measure latency)

## Run It

```bash
go run ./11_system_design_in_go/07_api_gateway_basics_mini
python3 ./11_system_design_in_go/07_api_gateway_basics_mini/main.py
```

## TL;DR

- API gateway = single entry point: routing + auth + rate limit + observability
- Middleware chain: auth -> rate limit -> request ID -> timeout -> route
- `context.WithTimeout` prevents hanging on slow downstream services
- Path-based routing dispatches to the correct service handler
- In production: use a dedicated gateway (Kong, Envoy) or cloud API gateway
- Gateway is the best place to add cross-cutting concerns (logging, metrics, auth)
- Always mention it is a single point of failure and needs high availability
