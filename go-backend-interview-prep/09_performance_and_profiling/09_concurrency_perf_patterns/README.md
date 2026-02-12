# Concurrency Performance Patterns

## What It Is

- **Worker pool**: fixed number of goroutines processing tasks from a channel -- bounded concurrency
- **ELI10:** Concurrency without limits is like opening all the browser tabs at once -- your machine begs for mercy.
- **Backpressure**: using buffered channels to limit how fast producers can push work

## Why It Matters

- Unbounded goroutine spawning can exhaust memory and kill your service
- **ELI10:** A worker pool is like having exactly four cashiers -- customers line up instead of everyone shouting at once.
- Interviewers expect you to know worker pool pattern and how to limit concurrency

## Syntax Cheat Sheet

```go
// Go: worker pool with buffered channel
jobs := make(chan int, 100)  // backpressure buffer
var wg sync.WaitGroup
for w := 0; w < numWorkers; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for job := range jobs { process(job) }
    }()
}
```

```python
# Python: ThreadPoolExecutor
from concurrent.futures import ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=4) as pool:
    results = pool.map(process, tasks)
```

> **Go**: goroutines are cheap (~2KB stack) but still need bounds.
> **Python**: threads are expensive (~8MB stack); always use a pool.

## Tiny Example

- `main.go` -- unbounded goroutines vs worker pool, timed comparison
- `main.py` -- unbounded threads vs ThreadPoolExecutor, timed comparison

## Common Interview Traps

- **Spawning a goroutine per request**: works at low scale, crashes at high scale
- **No backpressure**: producers flood consumers, memory grows unbounded
- **Forgetting WaitGroup**: program exits before workers finish
- **Channel deadlock**: unbuffered channel blocks if no receiver is ready
- **Over-sizing the pool**: too many workers can contend on shared resources

## What to Say in Interviews

- "I use a fixed worker pool with a buffered channel for backpressure"
- "Unbounded goroutines risk OOM -- the pool bounds memory and CPU usage"
- "The buffer size controls backpressure: producer blocks when the buffer is full"

## Run It

```bash
go run ./09_performance_and_profiling/09_concurrency_perf_patterns/
python ./09_performance_and_profiling/09_concurrency_perf_patterns/main.py
```

## TL;DR (Interview Summary)

- Worker pool: N goroutines reading from a shared channel
- Buffered channel = backpressure (producer blocks when full)
- Unbounded goroutines: fine for small N, dangerous at scale
- `sync.WaitGroup` to wait for all workers to finish
- Close the jobs channel to signal workers to exit
- Python: `ThreadPoolExecutor(max_workers=N)` is the equivalent
- Right pool size depends on workload: CPU-bound = numCPU, IO-bound = higher
