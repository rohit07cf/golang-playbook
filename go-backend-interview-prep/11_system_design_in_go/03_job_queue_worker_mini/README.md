# 03 -- Job Queue Worker (Mini)

## What We Are Building

- A producer that pushes jobs into a queue
- A worker pool that consumes and processes jobs concurrently with backpressure

## Requirements

**Functional:**
- Producer creates jobs with ID, payload, and priority
- Workers pick up jobs and process them (simulated work)
- Print job status: queued, processing, completed

**Non-functional:**
- Bounded queue -- backpressure when queue is full (producer blocks)
- Configurable worker count
- Graceful shutdown -- finish in-flight jobs before exit

## High-Level Design

```
Producer --> [buffered channel (bounded)] --> Worker 1
                                          --> Worker 2
                                          --> Worker 3
```

```
+----------+     +----------------+     +---------+
| Producer | --> | Buffered Chan  | --> | Worker  |
| (loop)   |     | (capacity = N) |     | Pool    |
+----------+     +----------------+     +---------+
                  backpressure               |
                  when full            WaitGroup
                                       for shutdown
```

- Buffered channel = bounded queue
- Channel full = producer blocks (backpressure)
- Close channel = signal workers to drain and exit

## Key Go Building Blocks Used

- `chan Job` (buffered) -- bounded job queue
- `sync.WaitGroup` -- wait for all workers to finish
- `time.Sleep` -- simulate work
- Goroutines -- worker pool
- Channel close -- signal completion

## Trade-Offs

- **Buffered channel vs external queue** -- channel is in-process; production uses Redis, RabbitMQ, SQS
- **Backpressure via blocking** -- simple but producer is stuck; alternative: drop or return error
- **No persistence** -- jobs lost on crash; production uses durable queues
- **No retry** -- failed jobs are gone; production adds dead-letter queue
- **Fixed worker count** -- no auto-scaling; production adjusts based on queue depth

## Common Interview Traps

- Forgetting to close the channel (workers hang forever)
- Not using WaitGroup (main exits before workers finish)
- Unbounded queue (channel with no buffer) -- OOM risk in production
- Confusing fan-out (one producer, many consumers) with fan-in (many producers, one consumer)
- Not mentioning dead-letter queues for failed jobs

## Run It

```bash
go run ./11_system_design_in_go/03_job_queue_worker_mini
python3 ./11_system_design_in_go/03_job_queue_worker_mini/main.py
```

## TL;DR

- Buffered channel = bounded queue with built-in backpressure
- Worker pool = N goroutines reading from the same channel
- Close channel to signal "no more jobs" -- workers drain and exit
- `sync.WaitGroup` ensures main waits for all workers to finish
- In production: use durable queue (Redis, SQS), add retries, dead-letter queue
- Mention backpressure explicitly in interviews -- it shows depth
