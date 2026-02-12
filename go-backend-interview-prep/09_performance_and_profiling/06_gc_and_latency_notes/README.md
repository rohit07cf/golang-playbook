# GC and Latency Notes

## What It Is

- **Go's GC**: concurrent, tri-color mark-and-sweep with very short stop-the-world (STW) pauses
- **ELI10:** The garbage collector is the office janitor -- works quietly, but if there is too much trash, everyone notices the delay.
- **Python's GC**: reference counting (immediate) + cyclic garbage collector (periodic)

## Why It Matters

- GC pauses can cause latency spikes in real-time services
- **ELI10:** If your janitor has to stop everyone to clean up, your service freezes -- reducing trash means fewer interruptions.
- Interviewers ask "how does Go's GC work?" and "how do you reduce GC pressure?"

## Syntax Cheat Sheet

```go
// Go: force GC and check stats
runtime.GC()
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Println("NumGC:", m.NumGC)
fmt.Println("PauseTotalNs:", m.PauseTotalNs)
// Tune GC: GOGC=100 (default), GOGC=200 (less frequent)
```

```python
# Python: gc module
import gc
gc.collect()                    # force cyclic GC
print(gc.get_stats())           # generation stats
print(sys.getrefcount(obj))     # reference count
```

> **Go**: concurrent GC, ~sub-ms STW pauses, tunable via `GOGC`.
> **Python**: refcount (instant free) + cyclic GC (periodic, can pause).

## Tiny Example

- `main.go` -- allocation loop showing GC counts and pause times
- `main.py` -- allocation loop showing gc.collect() stats and reference counting

## Common Interview Traps

- **Saying Go has stop-the-world GC**: it's mostly concurrent; STW phases are sub-millisecond
- **Not knowing GOGC**: `GOGC=100` means GC runs when heap doubles; increase to reduce frequency
- **Ignoring allocation rate**: GC frequency is proportional to allocation rate, not heap size alone
- **Thinking Python has no GC**: it does -- cyclic GC handles reference cycles
- **Confusing GC pause with total GC time**: pause is wall-clock delay; total CPU time is different

## What to Say in Interviews

- "Go's GC is concurrent mark-and-sweep with sub-millisecond STW pauses"
- "I reduce GC pressure by reducing allocations -- preallocate, use value types, sync.Pool"
- "GOGC controls GC frequency: higher values trade memory for fewer collections"

## Run It

```bash
go run ./09_performance_and_profiling/06_gc_and_latency_notes/
python ./09_performance_and_profiling/06_gc_and_latency_notes/main.py
```

## TL;DR (Interview Summary)

- Go GC: concurrent, tri-color mark-and-sweep, sub-ms STW pauses
- `GOGC=100` (default): GC when heap grows to 2x live data
- Reduce GC pressure: fewer allocs, preallocate, `sync.Pool`, value types
- `runtime.ReadMemStats` -- NumGC, PauseTotalNs, HeapAlloc
- Python: refcount (instant free) + cyclic GC (handles cycles)
- Allocation rate drives GC frequency -- not just heap size
- `GOMEMLIMIT` (Go 1.19+): soft memory limit for better GC behavior
