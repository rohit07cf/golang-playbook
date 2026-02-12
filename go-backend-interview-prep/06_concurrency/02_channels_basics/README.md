# Channels Basics

## What It Is

- A **channel** is a typed conduit for sending and receiving values between goroutines
- **ELI10:** A channel is a conveyor belt in a factory -- workers don't shout across the room, they drop boxes on the belt
- Created with `make(chan T)` -- unbuffered by default (send blocks until receiver is ready)

## Why It Matters

- Channels are Go's primary synchronization primitive: "share memory by communicating"
- Interviewers expect you to explain unbuffered vs buffered, close semantics, and range

## Syntax Cheat Sheet

```go
// Go: make, send, receive, close, range
ch := make(chan int)
go func() { ch <- 42 }()  // send
val := <-ch                 // receive (blocks)
close(ch)                   // close when done sending
for v := range ch { ... }   // iterate until closed
```

```python
# Python: queue.Queue as channel equivalent
import queue, threading
q = queue.Queue()
threading.Thread(target=lambda: q.put(42)).start()
val = q.get()  # blocks until item available
# No built-in "close" -- use sentinel values
```

> **Python differs**: `queue.Queue` has no close/range. Use a sentinel
> (like `None`) to signal "done". No compile-time type safety.

## Tiny Example

- `main.go` -- unbuffered channel send/receive, close + range, comma-ok receive
- `main.py` -- `queue.Queue` with sentinel-based close pattern

## Common Interview Traps

- **Send on closed channel panics**: only the sender should close
- **ELI10:** Sending to a channel with no receiver is like leaving a package at a door that nobody opens -- you'll wait forever
- **Receive on closed channel returns zero value**: use comma-ok `v, ok := <-ch`
- **Unbuffered = synchronous**: send blocks until someone receives
- **Forgetting to close**: `range ch` blocks forever if channel isn't closed
- **Multiple closers**: closing a channel twice panics -- only one goroutine should close

## What to Say in Interviews

- "Channels enforce synchronization at the type level -- send and receive are blocking operations"
- "I close channels from the sender side only, and receivers use range or comma-ok"
- "For simple value passing I use unbuffered; for decoupling producer/consumer I buffer"

## Run It

```bash
go run ./06_concurrency/02_channels_basics/
python ./06_concurrency/02_channels_basics/main.py
```

## TL;DR (Interview Summary)

- `ch := make(chan T)` -- unbuffered, send blocks until receive
- `ch <- val` sends; `val := <-ch` receives
- `close(ch)` -- only the sender closes; receiver gets zero + false
- `for v := range ch` iterates until channel is closed
- Send on closed channel = **panic**
- Receive on closed channel = zero value (use comma-ok to detect)
- Python: `queue.Queue` + sentinel for close semantics
