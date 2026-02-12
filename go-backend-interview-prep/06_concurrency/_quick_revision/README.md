# Concurrency -- Quick Revision

> One-screen cheat sheet. Skim before interviews.

---

## 1. Goroutines vs Threads vs Asyncio

```go
// Go: goroutine (~2 KB, M:N scheduled)
go func() { fmt.Println("hello") }()
```

```python
# Thread: OS thread (~1 MB, GIL limits CPU parallelism)
threading.Thread(target=func).start()
# asyncio: cooperative, single-threaded, good for I/O
asyncio.create_task(coro())
```

## 2. Channels: Unbuffered / Buffered + Close

```go
ch := make(chan int)      // unbuffered: send blocks until receive
ch := make(chan int, 5)   // buffered: send blocks when full
close(ch)                 // only sender closes
for v := range ch { }     // iterate until closed
```

```python
q = queue.Queue()           # unbounded by default
q = queue.Queue(maxsize=5)  # bounded
q.put(None)                 # sentinel = "close"
```

## 3. Select: Multiplexing

```go
select {
case v := <-ch1:             // first ready wins
case ch2 <- val:             // can also select on sends
case <-time.After(1*time.Second): // timeout
default:                     // non-blocking
}
```

```python
# asyncio: wait on multiple tasks
done, _ = await asyncio.wait(tasks, return_when=FIRST_COMPLETED)
# timeout
await asyncio.wait_for(coro(), timeout=1.0)
```

## 4. WaitGroup vs Mutex

```go
// WaitGroup: wait for goroutines to finish
var wg sync.WaitGroup
wg.Add(1); go func() { defer wg.Done(); work() }()
wg.Wait()

// Mutex: protect shared state
var mu sync.Mutex
mu.Lock(); counter++; mu.Unlock()
```

```python
# join = WaitGroup
for t in threads: t.join()
# Lock = Mutex
with lock: counter += 1
```

## 5. Worker Pool

```go
for w := 0; w < 3; w++ {
    go func() {
        for job := range jobs { results <- process(job) }
    }()
}
```

```python
with ThreadPoolExecutor(max_workers=3) as pool:
    futures = [pool.submit(process, j) for j in jobs]
```

## 6. Fan-Out / Fan-In

```go
// Fan-out: N workers read from one channel
// Fan-in:  merge N channels into one
merged := fanIn(worker1Out, worker2Out, worker3Out)
```

```python
workers = [asyncio.create_task(work(q)) for _ in range(3)]
await asyncio.gather(*workers)
```

## 7. Context Cancellation

```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
select {
case <-ctx.Done(): return ctx.Err()
case r := <-ch:    return r
}
```

```python
# asyncio
await asyncio.wait_for(coro(), timeout=5.0)
# threading
stop = threading.Event()
if stop.is_set(): return
```

## 8. Rate Limiting

```go
limiter := time.NewTicker(100 * time.Millisecond)
defer limiter.Stop()
for req := range requests { <-limiter.C; handle(req) }
```

```python
for req in requests:
    time.sleep(0.1)
    handle(req)
```

---

## Interview One-Liners

1. "Goroutines are ~2 KB user-space threads; Go can run millions"
2. "Channels are typed, synchronized pipes -- unbuffered = sync, buffered = async up to cap"
3. "select multiplexes channels; if multiple ready, picks randomly"
4. "WaitGroup waits for completion; mutex protects shared state"
5. "Worker pool = fixed goroutines + shared jobs channel"
6. "Context propagates cancellation down the call tree -- always defer cancel"
7. "Fan-out parallelizes; fan-in merges -- each stage owns its output channel"
8. "Always run `go test -race` to catch data races"

---

## TL;DR

- Goroutines = cheap (~2 KB); threads = expensive (~1 MB)
- Channels: unbuffered = sync, buffered = async, close = signal done
- `select` = multiplex channels, timeout, non-blocking with default
- `sync.WaitGroup` = wait for all; `sync.Mutex` = protect shared state
- Worker pool = bounded concurrency (N goroutines + jobs channel)
- `context.Context` = cancellation + timeout propagation
- Fan-out/fan-in = parallel pipeline stages
- Python: threads (GIL), asyncio (cooperative), queue.Queue (channel analog)
