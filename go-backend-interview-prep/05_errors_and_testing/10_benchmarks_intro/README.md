# Benchmarks Intro

## What It Is

- Go benchmarks use `testing.B` and run with `go test -bench=.`
- **ELI10:** A benchmark is a stopwatch for your code -- run it a million times and see who's fastest.
- The framework adjusts `b.N` automatically to get stable timing

## Why It Matters

- Performance questions come up in system design and backend interviews
- Knowing how to benchmark proves you can back up optimization claims with data

## Syntax Cheat Sheet

```go
// Go: benchmark function
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fib(20)
    }
}
// Run: go test -bench=. -benchmem
```

```python
# Python: timeit (stdlib analogy)
import timeit
t = timeit.timeit(lambda: fib(20), number=1000)
print(f"{t:.4f}s for 1000 runs")
```

## Tiny Example

- `main.go` -- defines `Fib` (recursive) and `FibIter` (iterative)
- `bench_test.go` -- benchmarks both approaches, shows `b.ResetTimer`
- `main.py` -- same two fibonacci functions
- `bench_test.py` -- uses `timeit` to compare both

## Common Interview Traps

- **Compiler optimizing away the result**: assign to a package-level variable to prevent it
- **Not using `b.N`**: the loop count must be `b.N`, not a hardcoded number
- **Benchmarking with `-count=1`**: run multiple times (`-count=5`) for stable results
- **Ignoring `-benchmem`**: memory allocations matter as much as speed
- **Micro-benchmarking without context**: real performance depends on the whole system

## What to Say in Interviews

- "I use `go test -bench=. -benchmem` to measure both time and allocations"
- "The framework adjusts `b.N` to get statistically stable results"
- "I assign benchmark results to a package-level var to prevent compiler optimization"

## Run It

```bash
go run ./05_errors_and_testing/10_benchmarks_intro/
python ./05_errors_and_testing/10_benchmarks_intro/main.py
go test -bench=. -benchmem ./05_errors_and_testing/10_benchmarks_intro/
python ./05_errors_and_testing/10_benchmarks_intro/bench_test.py
```

## TL;DR (Interview Summary)

- Benchmark functions: `BenchmarkXxx(b *testing.B)`
- Loop body runs `b.N` times (framework-controlled)
- Run: `go test -bench=. -benchmem`
- Use `b.ResetTimer()` after expensive setup
- Assign results to package-level var to prevent compiler elimination
- Python analogy: `timeit.timeit(func, number=N)`
