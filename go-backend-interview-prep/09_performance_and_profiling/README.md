# 09 -- Performance and Profiling

Measuring and improving Go program performance.

---

## Planned Topics

- [ ] Benchmarks -- `testing.B`, writing effective benchmarks
- [ ] pprof -- CPU profiling, memory profiling
- [ ] Escape analysis -- stack vs heap, `go build -gcflags="-m"`
- [ ] Memory allocation -- reducing allocs, pooling with `sync.Pool`
- [ ] String building -- `strings.Builder` vs concatenation
- [ ] Preallocation -- `make([]T, 0, n)` for slices, map hints
- [ ] Inlining -- what helps the compiler optimize
