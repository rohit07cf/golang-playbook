# Performance & Profiling -- Quick Revision

> One-screen refresher. Skim before your interview.

---

## When to Benchmark vs Profile

- **Benchmark**: "how fast is this function?" -> `go test -bench=. -benchmem`
- **Profile**: "where does the time go?" -> `runtime/pprof` + `go tool pprof`
- Benchmark first, profile only when you find something slow

## Key pprof Commands

```
go tool pprof cpu.prof          # open interactive shell
  top10                          # top 10 functions by flat time
  top -cum                       # sort by cumulative time
  list funcName                  # annotated source
  web                            # SVG call graph in browser
  peek funcName                  # callers + callees
  quit                           # exit
```

## Python cProfile Quick Commands

```bash
python -m cProfile -s cumulative script.py     # sort by cumtime
python -m cProfile -o output.prof script.py    # save to file
# Then in Python:
import pstats; pstats.Stats('output.prof').sort_stats('cumulative').print_stats(10)
```

## Allocation Reduction Rules

1. `make([]T, 0, n)` when size is known
2. `strings.Builder` + `Grow(n)` instead of `+=`
3. Return values, not pointers, in hot paths
4. `sync.Pool` for frequently allocated objects
5. `strconv.Itoa` instead of `fmt.Sprintf` in hot paths
6. Close `resp.Body` to return connections to pool

## 10 Tiny Go Snippets

```go
// 1. Benchmark function
func BenchmarkX(b *testing.B) {
    for i := 0; i < b.N; i++ { doWork() }
}
// go test -bench=. -benchmem
```

```go
// 2. CPU profile
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

```go
// 3. Memory stats
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Println(m.HeapAlloc, m.NumGC)
```

```go
// 4. Escape analysis check
// go build -gcflags="-m" ./pkg/
```

```go
// 5. Preallocate slice
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ { s = append(s, i) }
```

```go
// 6. strings.Builder
var b strings.Builder
b.Grow(n)
for i := 0; i < n; i++ { b.WriteString("x") }
s := b.String()
```

```go
// 7. Worker pool
jobs := make(chan int, 100)
for w := 0; w < numCPU; w++ {
    go func() { for j := range jobs { process(j) } }()
}
```

```go
// 8. Reused HTTP client
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{MaxIdleConnsPerHost: 10},
}
```

```go
// 9. Prevent dead code in benchmarks
var result string
func BenchmarkX(b *testing.B) {
    for i := 0; i < b.N; i++ { result = compute() }
}
```

```go
// 10. GC tuning
// GOGC=200 go run main.go    (less frequent GC)
// GOMEMLIMIT=512MiB           (soft memory limit, Go 1.19+)
```

## 10 Tiny Python Snippets

```python
# 1. timeit benchmark
import timeit
t = timeit.timeit(lambda: do_work(), number=10000)
```

```python
# 2. cProfile
import cProfile
cProfile.run('do_work()', sort='cumulative')
```

```python
# 3. tracemalloc
import tracemalloc
tracemalloc.start()
do_work()
snap = tracemalloc.take_snapshot()
for s in snap.statistics('lineno')[:5]: print(s)
```

```python
# 4. Reference counting
import sys
print(sys.getrefcount(obj))  # 2 = obj + arg
```

```python
# 5. List preallocation
s = [0] * n
for i in range(n): s[i] = compute(i)
```

```python
# 6. String join (fast)
parts = [f"item{i}" for i in range(n)]
result = ",".join(parts)
```

```python
# 7. ThreadPoolExecutor
from concurrent.futures import ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=4) as pool:
    results = list(pool.map(process, tasks))
```

```python
# 8. HTTP with timeout
import urllib.request
resp = urllib.request.urlopen(url, timeout=10)
```

```python
# 9. Queue backpressure
import queue
q = queue.Queue(maxsize=10)  # blocks when full
```

```python
# 10. GC control
import gc
gc.collect()          # force cyclic GC
gc.set_threshold(700, 10, 10)  # tune generations
```

## 10 Interview One-Liners

| # | Topic | One-Liner |
|---|-------|-----------|
| 1 | Benchmarks | `go test -bench=. -benchmem` -- ns/op + allocs/op in one command |
| 2 | pprof | `go tool pprof cpu.prof` then `top10` to find the hotspot |
| 3 | Memory | `runtime.ReadMemStats` for HeapAlloc, NumGC, PauseTotalNs |
| 4 | Escape | `go build -gcflags="-m"` -- returned pointers escape to heap |
| 5 | GC | Concurrent mark-and-sweep, sub-ms STW, tunable via GOGC |
| 6 | Strings | `strings.Builder` + `Grow(n)` -- O(1) allocs vs O(n) with += |
| 7 | Slices | `make([]T, 0, n)` -- preallocate to avoid repeated copying |
| 8 | Concurrency | Worker pool + buffered channel = bounded concurrency + backpressure |
| 9 | HTTP | Reuse `http.Client`, close `resp.Body`, set `Timeout` |
| 10 | Measure | Don't guess -- benchmark, then profile, then optimize the hotspot |

## TL;DR

- **Benchmark**: `go test -bench=. -benchmem` -- know ns/op and allocs/op
- **Profile**: `pprof.StartCPUProfile` -> `go tool pprof` -> `top10`
- **Memory**: `ReadMemStats` for runtime stats, `-memprofile` for per-function
- **Escape**: `go build -gcflags="-m"` -- pointers and interfaces cause heap allocs
- **GC**: concurrent, sub-ms pauses, reduce allocs = reduce GC pressure
- **Strings**: `strings.Builder` >> `+=` concat in loops
- **Slices/Maps**: preallocate with `make(T, 0, n)` when size is known
- **HTTP**: one `http.Client` per service, close body, set timeouts
- **Concurrency**: worker pool pattern, buffered channels for backpressure
- **Golden rule**: measure first, optimize second, verify with benchmarks
