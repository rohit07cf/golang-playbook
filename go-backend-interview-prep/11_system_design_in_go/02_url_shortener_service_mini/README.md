# 02 -- URL Shortener Service (Mini)

## What We Are Building

- An HTTP service that shortens URLs and resolves short codes back to the original
- In-memory storage with basic collision handling

## Requirements

**Functional:**
- `POST /shorten` with `{"url": "https://example.com"}` returns `{"code": "abc123"}`
- `GET /{code}` returns the original URL (or redirects)
- Short codes are 6 characters, base62 encoded

**Non-functional:**
- Concurrency-safe reads and writes
- Low collision probability
- Fast lookups (O(1) via map)

## High-Level Design

```
Client --> POST /shorten --> generate code --> store in map --> return code
Client --> GET /{code}   --> lookup map    --> return URL / 404
```

```
+----------+       +-----------+       +-------------+
|  Client  | ----> |  Handler  | ----> |  URLStore   |
+----------+       +-----------+       | (interface) |
                                       +-------------+
                                            |
                                       +-------------+
                                       | MemoryStore |
                                       | map[code]url|
                                       +-------------+
```

## Key Go Building Blocks Used

- `net/http` -- HTTP server + routing
- `sync.RWMutex` -- concurrent read/write safety on the URL map
- `crypto/rand` -- random code generation
- `encoding/json` -- request/response marshaling
- Interface for store -- easy to swap to DB later

## Trade-Offs

- **In-memory** -- fast but loses data on restart; mention Redis or DB for persistence
- **Random codes vs hash-based** -- random avoids needing the URL as input but needs collision check
- **6-char base62** -- ~56 billion combinations, good for demo; production uses 7-8 chars
- **No expiration** -- production would add TTL for cleanup
- **Single instance** -- no distributed ID generation; mention snowflake IDs for scale

## Common Interview Traps

- Not explaining your encoding scheme (base62 vs base64 vs hash truncation)
- Forgetting collision handling -- what if two URLs get the same code?
- Using MD5/SHA and calling it "unique" -- hash collisions exist
- Not mentioning read-heavy ratio (reads >> writes in URL shorteners)
- Ignoring analytics (click tracking is a common follow-up question)

## Run It

```bash
go run ./11_system_design_in_go/02_url_shortener_service_mini
python3 ./11_system_design_in_go/02_url_shortener_service_mini/main.py
```

## TL;DR

- Generate random base62 codes, check for collision, retry if needed
- `sync.RWMutex` allows concurrent reads, exclusive writes
- Interface for store layer makes it testable and swappable
- Mention persistence (Redis/DB), analytics, and expiration in interviews
- Read-heavy workload: optimize for fast lookups
- Short code length = trade-off between URL size and collision probability
