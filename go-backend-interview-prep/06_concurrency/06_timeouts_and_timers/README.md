# Timeouts and Timers

## What It Is

- `time.After(d)` returns a channel that sends after duration `d` (one-shot)
- `time.NewTimer(d)` is a reusable one-shot; `time.NewTicker(d)` fires repeatedly

## Why It Matters

- Timeouts prevent goroutines from blocking forever on slow operations
- Interviewers test timeout patterns with `select` and `time.After`

## Syntax Cheat Sheet

```go
// Go: timeout via select
select {
case v := <-ch:              // happy path
case <-time.After(2*time.Second): // timeout
}
// Ticker fires repeatedly
ticker := time.NewTicker(500 * time.Millisecond)
defer ticker.Stop()
```

```python
# Python (asyncio): wait_for with timeout
result = await asyncio.wait_for(coro(), timeout=2.0)
# Python (threading): queue.get(timeout=2.0)
```

> **Python differs**: `asyncio.wait_for` raises `TimeoutError`.
> `queue.Queue.get(timeout=N)` raises `queue.Empty`. No channel-based
> timer equivalent.

## Tiny Example

- `main.go` -- `time.After` timeout in select, `NewTimer` reset, `NewTicker` loop with stop
- `main.py` -- asyncio `wait_for` timeout, periodic loop with `asyncio.sleep`

## Common Interview Traps

- **time.After leaks in loops**: each call creates a new timer; use `NewTimer` + `Reset` in loops
- **Forgetting ticker.Stop()**: leaked tickers keep firing, wasting resources
- **Timer.Reset race**: only reset a timer after it has fired or been stopped + drained
- **time.Sleep vs select+After**: Sleep blocks the goroutine; select+After is cancellable

## What to Say in Interviews

- "I use select + time.After for simple timeouts and NewTicker for periodic work"
- "In loops, I use NewTimer + Reset instead of time.After to avoid allocations"
- "I always defer ticker.Stop() to prevent resource leaks"

## Run It

```bash
go run ./06_concurrency/06_timeouts_and_timers/
python ./06_concurrency/06_timeouts_and_timers/main.py
```

## TL;DR (Interview Summary)

- `time.After(d)` -- one-shot channel timer, great for select timeout
- `time.NewTimer(d)` -- reusable one-shot, call `Reset` / `Stop`
- `time.NewTicker(d)` -- repeated ticks, always `defer Stop()`
- Don't use `time.After` in loops (leaks timers) -- use `NewTimer` + `Reset`
- `time.Sleep` blocks the goroutine; select + timer is cancellable
- Python: `asyncio.wait_for(timeout=)`, `asyncio.sleep` for periodic
