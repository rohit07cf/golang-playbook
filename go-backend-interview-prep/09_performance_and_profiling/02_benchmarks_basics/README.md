# Benchmarks Basics

## What It Is

- **Go benchmarks**: functions starting with `Benchmark` in `_test.go` files, using `testing.B`
- **ELI10:** A benchmark is a stopwatch race for your functions -- run them a million times and crown the winner.
- **Python timeit**: stdlib module that runs a statement many times and reports average time

## Why It Matters

- Interviewers expect you to know how to measure performance properly (not with `time.Now()`)
- **ELI10:** Without benchmarks, "I think it's faster" is just a guess -- benchmarks turn opinions into facts.
- `-benchmem` reveals allocations-per-op -- the metric that matters most for Go optimization

## Syntax Cheat Sheet

```go
// Go: benchmark function
func BenchmarkFoo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        foo()
    }
}
// Run: go test -bench=. -benchmem
```

```python
# Python: timeit
import timeit
t = timeit.timeit(lambda: foo(), number=10000)
print(f"{t/10000*1e6:.1f} us/op")
```

> **Go differs**: the framework auto-adjusts `b.N` for stable results.
> Python's `timeit` requires you to choose `number` manually.

## Tiny Example

- `main.go` -- demonstrates what benchmarks test (concat vs builder); run `bench_test.go` for real results
- `main.py` -- uses `timeit` to compare string approaches; `bench_test.py` has structured benchmarks
- `bench_test.go` -- real Go benchmarks with `BenchmarkConcat` and `BenchmarkBuilder`
- `bench_test.py` -- Python benchmark script using `timeit`

## Common Interview Traps

- **Using `time.Now()` for benchmarks**: doesn't account for warmup, GC pauses, or compiler optimization
- **Forgetting `-benchmem`**: without it you miss allocation counts
- **Dead code elimination**: if you don't use the result, the compiler may skip the work
- **Benchmarking too little**: `b.N` auto-scales, but custom loops may not run enough iterations
- **Comparing Go and Python benchmarks directly**: different runtimes, different overhead

## What to Say in Interviews

- "I use `go test -bench=. -benchmem` to get ns/op and allocs/op in one pass"
- "The framework adjusts b.N automatically so results are stable across runs"
- "I always check allocs/op -- reducing allocations often matters more than faster algorithms"

## Run It

```bash
# Run the demo
go run ./09_performance_and_profiling/02_benchmarks_basics/
python ./09_performance_and_profiling/02_benchmarks_basics/main.py

# Run actual benchmarks
go test -bench=. -benchmem ./09_performance_and_profiling/02_benchmarks_basics/
python ./09_performance_and_profiling/02_benchmarks_basics/bench_test.py
```

## TL;DR (Interview Summary)

- `go test -bench=. -benchmem` -- the one command you must know
- Output: `BenchmarkName-CPU  N  ns/op  B/op  allocs/op`
- `b.N` auto-adjusts -- don't set iteration count manually
- Assign results to a package-level var to prevent dead code elimination
- `-benchmem` shows bytes and allocations per operation
- Python: `timeit.timeit(func, number=N)` -- manual iteration count
- Always benchmark the specific operation, not the setup
