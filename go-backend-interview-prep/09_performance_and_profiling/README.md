# 09 -- Performance and Profiling

## What This Module Covers

Go is fast by default, but writing performant code still requires understanding
where time and memory go. This module teaches you to measure, profile, and
reason about performance -- the skills interviewers test when they ask
"how would you optimize this?"

You'll learn to benchmark, profile CPU and memory, understand escape analysis
and GC, and avoid the most common performance pitfalls in strings, slices,
maps, concurrency, and HTTP clients.

Every example includes `main.go` + Python equivalent `main.py`.

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `big_o_and_hotspots` | Algorithmic complexity, finding the slow part |
| 02 | `benchmarks_basics` | `testing.B`, `-benchmem`, Python `timeit` |
| 03 | `cpu_profiling_pprof_intro` | `runtime/pprof`, `go tool pprof`, `cProfile` |
| 04 | `memory_profiling_and_allocations` | Heap allocations, `ReadMemStats`, `tracemalloc` |
| 05 | `escape_analysis_and_pointers` | Stack vs heap, `-gcflags="-m"`, Python's heap model |
| 06 | `gc_and_latency_notes` | GC pacing, STW pauses, Python ref counting |
| 07 | `strings_bytes_and_builders` | `strings.Builder`, `[]byte`, `''.join()` |
| 08 | `maps_slices_perf_gotchas` | Preallocation, map growth, key costs |
| 09 | `concurrency_perf_patterns` | Worker pools, backpressure, bounded concurrency |
| 10 | `http_perf_basics` | Keep-alive, connection reuse, client timeouts |

---

## 10-Min Revision Path

1. Skim `01_big_o_and_hotspots` -- O(n^2) vs O(n) with timing proof
2. Review `02_benchmarks_basics` -- `go test -bench=. -benchmem` output format
3. Read `03_cpu_profiling_pprof_intro` -- pprof commands you'll quote in interviews
4. Scan `04_memory_profiling_and_allocations` -- allocation-per-op reasoning
5. Glance at `05_escape_analysis_and_pointers` -- `-gcflags="-m"` output
6. Skim `07_strings_bytes_and_builders` -- Builder vs concat benchmark
7. Read `09_concurrency_perf_patterns` -- worker pool vs unbounded goroutines
8. Finish with `_quick_revision/README.md` -- one-screen cheat sheet

---

## Common Performance Mistakes

- Concatenating strings in a loop instead of using `strings.Builder`
- Not preallocating slices when length is known (`make([]T, 0, n)`)
- Benchmarking with `time.Now()` instead of `testing.B` (wrong results due to warmup)
- Ignoring allocations-per-op in benchmark output
- Returning pointers from functions unnecessarily (forces heap allocation)
- Creating new `http.Client` per request (no connection reuse)
- Spawning unbounded goroutines without a worker pool
- Never profiling -- guessing where the bottleneck is instead of measuring
- Forgetting that map iteration order is random and map growth is expensive
- Using `fmt.Sprintf` in hot paths when `strconv` is faster

---

## TL;DR

- **Measure first**: `go test -bench=. -benchmem` before optimizing
- **Profile**: `runtime/pprof` for CPU + memory, `go tool pprof` to analyze
- **Escape analysis**: `go build -gcflags="-m"` shows what goes to heap
- **Strings**: `strings.Builder` for loops; avoid `+=` concat
- **Slices**: `make([]T, 0, n)` when you know the capacity
- **HTTP**: reuse `http.Client`, set timeouts, keep connections alive
- **Concurrency**: worker pool pattern, buffered channels for backpressure
