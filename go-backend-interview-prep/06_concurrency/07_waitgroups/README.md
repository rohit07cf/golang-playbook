# WaitGroups

## What It Is

- `sync.WaitGroup` blocks until a set of goroutines finish
- **ELI10:** WaitGroup is a headcount -- "I sent 5 people out, I'm not leaving until all 5 check back in"
- Three methods: `Add(n)`, `Done()`, `Wait()`

## Why It Matters

- The simplest way to wait for multiple goroutines -- used everywhere
- Interviewers expect you to know WaitGroup vs channels for synchronization

## Syntax Cheat Sheet

```go
// Go: WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
}()
wg.Wait()  // blocks until counter is 0
```

```python
# Python: thread.join() or concurrent.futures
threads = [threading.Thread(target=do_work) for _ in range(5)]
for t in threads: t.start()
for t in threads: t.join()  # wait for all
```

> **Python differs**: `thread.join()` is the direct equivalent.
> `concurrent.futures.wait()` also works for `ThreadPoolExecutor`.

## Tiny Example

- `main.go` -- launches workers with WaitGroup, shows Add before go, defer Done
- `main.py` -- same with thread.join() and concurrent.futures

## Common Interview Traps

- **Add before go**: call `wg.Add(1)` **before** launching the goroutine, not inside it
- **ELI10:** Calling wg.Add() inside the goroutine is like counting people after they already left the building
- **Passing WaitGroup by value**: always pass `*sync.WaitGroup` (pointer), not a copy
- **Negative counter**: calling `Done()` more than `Add()` panics
- **WaitGroup vs channel**: WaitGroup = "wait for all"; channel = "communicate results"
- **Forgetting defer Done()**: if the goroutine panics, Done never runs without defer

## What to Say in Interviews

- "I call `wg.Add(1)` before launching each goroutine and `defer wg.Done()` inside"
- "WaitGroup is for fan-out-wait-for-all; if I need results back, I combine it with channels"
- "I always pass WaitGroup by pointer to avoid copying the counter"

## Run It

```bash
go run ./06_concurrency/07_waitgroups/
python ./06_concurrency/07_waitgroups/main.py
```

## TL;DR (Interview Summary)

- `wg.Add(n)` before launching, `defer wg.Done()` inside, `wg.Wait()` to block
- Always call Add **before** `go` -- not inside the goroutine
- Always pass by pointer (`*sync.WaitGroup`)
- WaitGroup = wait for completion; channels = communicate data
- Calling Done more than Add = panic (negative counter)
- Python: `thread.join()` or `concurrent.futures.wait()`
