# Select

## What It Is

- `select` waits on **multiple channel operations** simultaneously
- Like a `switch` but for channels -- whichever is ready first wins

## Why It Matters

- `select` is how Go implements multiplexing, timeouts, and non-blocking channel ops
- Interviewers ask about `select` in nearly every concurrency question

## Syntax Cheat Sheet

```go
// Go: select on multiple channels
select {
case v := <-ch1:  fmt.Println("from ch1:", v)
case v := <-ch2:  fmt.Println("from ch2:", v)
case <-time.After(1 * time.Second): fmt.Println("timeout")
default:           fmt.Println("nothing ready")
}
```

```python
# Python: no direct select. Closest options:
# asyncio: await asyncio.wait(tasks, return_when=FIRST_COMPLETED)
# threading: poll multiple queues or use selectors
```

> **Python differs**: no native channel select. Use `asyncio.wait` for
> async tasks or manually poll queues for threads.

## Tiny Example

- `main.go` -- select on two channels, timeout with `time.After`, non-blocking with `default`
- `main.py` -- asyncio-based equivalent using `asyncio.wait` for multiplexing

## Common Interview Traps

- **Random selection**: if multiple cases are ready, Go picks one **randomly**
- **Default makes it non-blocking**: without `default`, select blocks until a case is ready
- **Nil channels are ignored**: a nil channel in select is never ready (useful for disabling cases)
- **Empty select blocks forever**: `select {}` is a permanent block (used in long-running servers)
- **Leaking goroutines via unused channel cases**: always provide a timeout or done channel

## What to Say in Interviews

- "select multiplexes channel operations -- it blocks until one case is ready, or falls through to default"
- "I use select with time.After for timeouts and with a done channel for cancellation"
- "If multiple cases are ready, Go selects one at random to avoid starvation"

## Run It

```bash
go run ./06_concurrency/05_select/
python ./06_concurrency/05_select/main.py
```

## TL;DR (Interview Summary)

- `select` waits on multiple channels, picks the first ready
- If multiple ready: **random** choice (prevents starvation)
- `default` case = non-blocking
- `time.After` in select = timeout pattern
- Nil channel = case is skipped (disables it)
- `select {}` blocks forever
- Python: `asyncio.wait` or manual queue polling
