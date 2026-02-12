# CPU Profiling / pprof Intro

## What It Is

- **pprof**: Go's built-in profiling tool -- collects CPU samples showing where time is spent
- **cProfile**: Python's stdlib profiler -- counts function calls and cumulative time

## Why It Matters

- Benchmarks tell you HOW FAST; profiling tells you WHERE the time goes
- Interviewers expect you to know the pprof workflow: collect, analyze, optimize

## Syntax Cheat Sheet

```go
// Go: collect CPU profile to file
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
// ... do work ...
// Analyze: go tool pprof cpu.prof
```

```python
# Python: cProfile
import cProfile, pstats
cProfile.run('do_work()', 'output.prof')
p = pstats.Stats('output.prof')
p.sort_stats('cumulative').print_stats(10)
```

> **Go**: sampling-based (low overhead, production-safe).
> **Python**: deterministic (counts every call, higher overhead).

## Tiny Example

- `main.go` -- writes a CPU profile to `cpu.prof` while running a workload; shows pprof commands
- `main.py` -- uses cProfile to profile the same workload; prints top functions

## Common Interview Traps

- **Profiling in production without sampling**: Go's pprof is safe; Python's cProfile is not
- **Forgetting to stop the profiler**: `defer pprof.StopCPUProfile()` -- or the file is incomplete
- **Reading pprof output wrong**: `flat` = time in this function; `cum` = time including callees
- **Not enough samples**: short workloads produce noisy profiles -- run longer
- **Confusing CPU and wall-clock time**: sleeping goroutines don't show in CPU profiles

## What to Say in Interviews

- "I collect a CPU profile with runtime/pprof, then analyze with go tool pprof"
- "I look at the top functions by flat time first, then drill into cumulative callers"
- "pprof is sampling-based and production-safe -- I can enable it in live services via net/http/pprof"

## Run It

```bash
# Go: generate and analyze CPU profile
go run ./09_performance_and_profiling/03_cpu_profiling_pprof_intro/
go tool pprof cpu.prof
# Inside pprof: top10, list funcName, web (generates SVG)

# Python: run with cProfile
python ./09_performance_and_profiling/03_cpu_profiling_pprof_intro/main.py
```

## TL;DR (Interview Summary)

- `runtime/pprof.StartCPUProfile(f)` + `defer StopCPUProfile()` -- collect profile
- `go tool pprof cpu.prof` -- interactive analysis
- Key pprof commands: `top`, `list`, `web`, `peek`
- `flat` = time in function; `cum` = time including callees
- `net/http/pprof` for live services (import for side effects)
- Python: `cProfile.run()` + `pstats.Stats().sort_stats('cumulative')`
- Profiling answers "where" -- benchmarks answer "how fast"
