# Fan-Out, Fan-In Pipeline

## What It Is

- **Fan-out**: multiple goroutines read from the same input channel (parallelism)
- **ELI10:** Fan-out is handing copies of the exam to 10 graders. Fan-in is collecting all their scores into one pile.
- **Fan-in**: multiple channels merge into one output channel
- **Pipeline**: chained stages where each stage's output is the next stage's input

## Why It Matters

- Pipelines are Go's answer to data processing -- like Unix pipes but typed and concurrent
- Interviewers frequently ask you to design or implement a multi-stage pipeline

## Syntax Cheat Sheet

```go
// Go: pipeline stage
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in { out <- v * v }
        close(out)
    }()
    return out
}
// Fan-in: merge multiple channels into one
```

```python
# Python: asyncio pipeline with queues
async def square(in_q, out_q):
    while (v := await in_q.get()) is not None:
        await out_q.put(v * v)
    await out_q.put(None)
```

> **Python differs**: use `asyncio.Queue` for async pipelines or
> `queue.Queue` with threads. No channel-based fan-in built-in.

## Tiny Example

- `main.go` -- 3-stage pipeline (generate -> square -> filter), fan-out with 3 workers, fan-in merge
- `main.py` -- same pipeline using asyncio tasks and queues

## Common Interview Traps

- **Forgetting to close output channels**: downstream `range` blocks forever
- **Fan-in ordering**: merged output is in completion order, not input order
- **Goroutine leaks**: if downstream stops reading, upstream goroutines block on send
- **Pipeline backpressure**: unbounded channels can consume unlimited memory
- **Error handling**: errors should flow through channels or context, not just be logged

## What to Say in Interviews

- "I structure data processing as a pipeline of stages connected by channels"
- "Fan-out parallelizes a stage; fan-in merges results back into one stream"
- "Each stage owns its output channel and closes it when done"

## Run It

```bash
go run ./06_concurrency/12_fanout_fanin_pipeline/
python ./06_concurrency/12_fanout_fanin_pipeline/main.py
```

## TL;DR (Interview Summary)

- **Pipeline**: generate -> stage1 -> stage2 -> consume, connected by channels
- **Fan-out**: N goroutines reading from one channel (parallel processing)
- **Fan-in**: merge N channels into one (use WaitGroup to know when to close)
- Each stage: receive from input, process, send to output, close output when done
- Results arrive in **completion order** (not input order) after fan-out
- Prevent leaks: always close output channels, use context for cancellation
- Python: asyncio queues or threading queues for the same pattern
