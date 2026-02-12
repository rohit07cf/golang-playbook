# Worker Pools

## What It Is

- A **worker pool** is a fixed number of goroutines reading jobs from a shared channel
- **ELI10:** A worker pool is 5 cashiers at a grocery store instead of spawning 500 random people at checkout
- Pattern: N workers consume from a jobs channel, send results to a results channel

## Why It Matters

- Bounds concurrency -- prevents launching a million goroutines at once
- **ELI10:** Without a pool, you'd hire a new cashier for every customer -- your store would run out of space
- This is one of the **most asked** Go concurrency patterns in interviews

## Syntax Cheat Sheet

```go
// Go: worker pool
jobs := make(chan Job, 100)
results := make(chan Result, 100)
for w := 0; w < numWorkers; w++ {
    go worker(jobs, results)
}
```

```python
# Python: ThreadPoolExecutor (stdlib)
from concurrent.futures import ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=3) as pool:
    futures = [pool.submit(work, job) for job in jobs]
```

> **Python differs**: `ThreadPoolExecutor` handles the pool internally.
> For manual control, use threads + `queue.Queue` (same pattern as Go).

## Tiny Example

- `main.go` -- 3 workers process 9 jobs, collect results via channel
- `main.py` -- same with `ThreadPoolExecutor` and manual thread+queue version

## Common Interview Traps

- **Forgetting to close the jobs channel**: workers range over jobs and block forever
- **Closing results too early**: close results only after all workers finish (use WaitGroup)
- **Unbounded goroutines**: without a pool, `go work(j)` in a loop = unbounded concurrency
- **Worker panics**: one panicking worker doesn't affect others (but you should recover)
- **Result ordering**: results come back in completion order, not submission order

## What to Say in Interviews

- "I create a fixed pool of workers that read from a shared jobs channel"
- "I use a WaitGroup to know when all workers are done, then close the results channel"
- "This bounds concurrency to N workers, preventing resource exhaustion"

## Run It

```bash
go run ./06_concurrency/09_worker_pools/
python ./06_concurrency/09_worker_pools/main.py
```

## TL;DR (Interview Summary)

- Pattern: N workers, 1 jobs channel, 1 results channel
- Workers: `for job := range jobs { results <- process(job) }`
- Close jobs channel when all jobs are sent
- Close results channel after all workers finish (WaitGroup)
- Results arrive in **completion order**, not submission order
- Bounds concurrency to N (prevents resource exhaustion)
- Python: `ThreadPoolExecutor` or manual threads + queue
