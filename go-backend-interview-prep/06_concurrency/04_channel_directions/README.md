# Channel Directions

## What It Is

- Go channels can be **directional**: `chan<- T` (send-only) or `<-chan T` (receive-only)
- **ELI10:** Channel directions are "one-way street" signs -- a send-only channel means no peeking at incoming traffic
- Functions declare which direction they need -- the compiler enforces it

## Why It Matters

- Directional channels make APIs self-documenting and prevent misuse at compile time
- Interviewers test whether you understand producer/consumer channel ownership

## Syntax Cheat Sheet

```go
// Go: directional channel types
func producer(out chan<- int) { out <- 42 }  // can only send
func consumer(in <-chan int)  { v := <-in }  // can only receive
// Bidirectional chan auto-converts to directional
```

```python
# Python: no directional queues -- convention only
def producer(out: queue.Queue) -> None:
    out.put(42)  # "send-only" by convention
def consumer(inp: queue.Queue) -> None:
    v = inp.get()  # "receive-only" by convention
```

> **Python differs**: no compile-time direction enforcement. You rely
> on naming conventions and documentation.

## Tiny Example

- `main.go` -- producer sends to `chan<-`, consumer reads from `<-chan`, shows compile-time safety
- `main.py` -- same pattern with `queue.Queue`, direction enforced by convention only

## Common Interview Traps

- **Can't close a receive-only channel**: `close(in)` on `<-chan` won't compile
- **Bidirectional converts to directional**: `ch := make(chan int)` passed to `chan<- int` is fine
- **Can't convert directional back to bidirectional**: one-way conversion only
- **Closing from wrong side**: convention is sender closes, directional types enforce this

## What to Say in Interviews

- "I use directional channel types to make function signatures self-documenting"
- "The compiler prevents sending on a receive-only channel or closing a receive-only channel"
- "Bidirectional channels implicitly convert to directional when passed to functions"

## Run It

```bash
go run ./06_concurrency/04_channel_directions/
python ./06_concurrency/04_channel_directions/main.py
```

## TL;DR (Interview Summary)

- `chan<- T` = send-only; `<-chan T` = receive-only
- Compiler enforces direction -- can't receive on send-only
- Bidirectional `chan T` auto-converts to either direction
- Can't convert directional back to bidirectional
- Only send-only side can close a channel
- Python has no equivalent -- convention only
