# Buffered Channels

## What It Is

- A **buffered channel** has a capacity: `make(chan T, size)`
- Sends only block when the buffer is **full**; receives block when it's **empty**

## Why It Matters

- Buffered channels decouple producer and consumer speed
- Interviewers ask "when would you use buffered vs unbuffered?"

## Syntax Cheat Sheet

```go
// Go: buffered channel
ch := make(chan int, 3)  // capacity 3
ch <- 1                   // does NOT block (buffer has room)
ch <- 2
ch <- 3
// ch <- 4              // WOULD block (buffer full)
```

```python
# Python: queue.Queue with maxsize
import queue
q = queue.Queue(maxsize=3)
q.put(1)  # does not block
q.put(2)
q.put(3)
# q.put(4) would block (queue full)
```

> **Python differs**: `queue.Queue(maxsize=N)` is the direct equivalent.
> `maxsize=0` means unlimited (unlike Go where 0 = unbuffered/synchronous).

## Tiny Example

- `main.go` -- sends multiple values into a buffered channel, drains them, shows blocking behavior
- `main.py` -- same pattern with `queue.Queue(maxsize=N)`

## Common Interview Traps

- **Buffered != async**: sends still block when full
- **Buffer size 0 = unbuffered**: `make(chan int, 0)` is the same as `make(chan int)`
- **Deadlock with single goroutine**: sending to a full unbuffered channel in main blocks forever
- **Overly large buffers hide bugs**: if you need a huge buffer, the design might be wrong
- **len(ch) and cap(ch)**: `len` = items in buffer, `cap` = buffer size

## What to Say in Interviews

- "I use unbuffered channels for synchronization and buffered for decoupling speed differences"
- "A buffered channel of size 1 can act as a simple semaphore or signal"
- "I check `len(ch)` for current items and `cap(ch)` for the buffer capacity"

## Run It

```bash
go run ./06_concurrency/03_buffered_channels/
python ./06_concurrency/03_buffered_channels/main.py
```

## TL;DR (Interview Summary)

- `make(chan T, N)` -- buffer size N; sends block only when full
- Unbuffered (`N=0`): send blocks until receive (synchronous handoff)
- Buffered (`N>0`): send succeeds immediately if buffer has room
- `len(ch)` = current items; `cap(ch)` = total capacity
- Buffer size 1 = simple signal / semaphore
- Don't over-buffer -- it hides design issues
- Python: `queue.Queue(maxsize=N)`
