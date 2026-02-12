# Memory Profiling and Allocations

## What It Is

- **Memory profiling**: measuring how much heap memory your program allocates and where
- **ELI10:** Memory profiling is checking who left all their stuff on the floor -- find the messy function, clean it up.
- **Allocations-per-op**: the key metric from `-benchmem` -- fewer allocs = less GC pressure

## Why It Matters

- Every heap allocation eventually triggers garbage collection, adding latency
- **ELI10:** Every allocation is trash the garbage collector has to pick up later -- less trash means less waiting.
- Interviewers ask "how would you reduce allocations?" -- you need to know where they happen

## Syntax Cheat Sheet

```go
// Go: read memory stats
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("HeapAlloc: %d KB\n", m.HeapAlloc/1024)
fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
fmt.Printf("NumGC: %d\n", m.NumGC)
```

```python
# Python: tracemalloc
import tracemalloc
tracemalloc.start()
# ... do work ...
snapshot = tracemalloc.take_snapshot()
for stat in snapshot.statistics('lineno')[:5]:
    print(stat)
```

> **Go**: `runtime.ReadMemStats` for global stats, pprof for per-function allocs.
> **Python**: `tracemalloc` traces per-line allocations.

## Tiny Example

- `main.go` -- compares allocation-heavy vs allocation-light approaches, prints MemStats
- `main.py` -- uses tracemalloc to show top allocation sites

## Common Interview Traps

- **Ignoring allocs/op in benchmark output**: bytes/op alone doesn't tell the full story
- **Not understanding TotalAlloc vs HeapAlloc**: TotalAlloc is cumulative; HeapAlloc is current live
- **Over-optimizing**: reducing allocs from 3 to 1 in a cold path doesn't matter
- **Forgetting sync.Pool exists**: reuse buffers for hot-path allocations
- **Not profiling before optimizing**: measure first, then reduce allocs in the hotspot

## What to Say in Interviews

- "I check allocs/op with -benchmem and focus on reducing allocations in the hot path"
- "I use runtime.ReadMemStats to get heap size and GC count at runtime"
- "For per-function allocation info, I use go tool pprof with a heap profile"

## Run It

```bash
go run ./09_performance_and_profiling/04_memory_profiling_and_allocations/
python ./09_performance_and_profiling/04_memory_profiling_and_allocations/main.py

# Go heap profile (alternative):
# go test -bench=. -benchmem -memprofile=mem.prof ./path/
# go tool pprof mem.prof
```

## TL;DR (Interview Summary)

- `runtime.ReadMemStats(&m)` -- HeapAlloc, TotalAlloc, NumGC
- `-benchmem` in benchmarks shows B/op and allocs/op
- `go tool pprof mem.prof` -- per-function allocation breakdown
- Reduce allocs in hot paths: preallocate, reuse buffers, avoid unnecessary pointers
- `sync.Pool` for frequently allocated/freed objects
- Python: `tracemalloc.start()` + `take_snapshot()` for per-line tracking
- Focus on allocs/op, not just bytes/op
