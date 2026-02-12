# 04 -- Cache Service (Mini)

## What We Are Building

- An in-memory cache with TTL (time-to-live) expiration
- Simple LRU-style eviction when cache reaches max capacity

## Requirements

**Functional:**
- `Set(key, value, ttl)` -- store a value with expiration
- `Get(key)` -- return value if present and not expired
- Evict expired entries on access and periodically
- Evict least-recently-used when capacity is exceeded

**Non-functional:**
- Concurrency-safe
- O(1) average for get and set
- Configurable max capacity and default TTL

## High-Level Design

```
Client --> Get(key) --> check map --> expired? evict : return
Client --> Set(key) --> capacity full? evict LRU --> store
```

```
+--------+       +------------------+
| Client | ----> |     Cache        |
+--------+       | map[key]*entry   |
                 | + doubly-linked  |
                 |   list for LRU   |
                 +------------------+
```

- Map gives O(1) lookup
- Linked list tracks access order (most recent at front)
- On access: move entry to front
- On eviction: remove from back (least recent)

## Key Go Building Blocks Used

- `sync.Mutex` -- protect shared cache state
- `container/list` -- doubly-linked list for LRU ordering
- `time.Now()` + `time.Duration` -- TTL expiration checks
- `time.NewTicker` -- periodic cleanup goroutine

## Trade-Offs

- **In-memory only** -- fast but limited by RAM; production uses Redis/Memcached
- **Simple LRU** -- good default; alternatives: LFU, ARC, random eviction
- **Global mutex** -- simple but limits throughput; sharded cache scales better
- **Eager + lazy eviction** -- lazy on access + periodic sweep is a good balance
- **No persistence** -- cache is warm only after population; cold start is slow

## Common Interview Traps

- Forgetting TTL -- cache without expiration serves stale data forever
- Not mentioning eviction policy -- what happens when cache is full?
- Using only lazy eviction -- expired entries consume memory until accessed
- Confusing cache-aside vs write-through vs write-behind patterns
- Not discussing cache stampede (thundering herd on cache miss)

## Run It

```bash
go run ./11_system_design_in_go/04_cache_service_mini
python3 ./11_system_design_in_go/04_cache_service_mini/main.py
```

## TL;DR

- Map + linked list = O(1) LRU cache
- TTL prevents serving stale data; check expiration on every `Get`
- Periodic cleanup goroutine sweeps expired entries
- `sync.Mutex` for concurrency safety
- In production: use Redis for shared cache across instances
- Mention eviction policy and cache stampede in interviews
