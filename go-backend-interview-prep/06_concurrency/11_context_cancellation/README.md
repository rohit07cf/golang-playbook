# Context Cancellation

## What It Is

- `context.Context` carries **deadlines, cancellation signals, and values** across API boundaries
- **ELI10:** Context is the "stop everything" button with a timer and a reason attached
- Created with `context.WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`

## Why It Matters

- Every production Go server uses context for request cancellation and timeouts
- **ELI10:** Without context, cancelling a goroutine tree is like trying to stop a rumor -- good luck
- Interviewers test whether you propagate context through function calls correctly

## Syntax Cheat Sheet

```go
// Go: context cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Check in goroutine
select {
case <-ctx.Done(): return ctx.Err()
case result := <-ch: return result
}
```

```python
# Python (asyncio): task cancellation
task = asyncio.create_task(work())
task.cancel()  # sends CancelledError

# Python (threading): Event flag
stop = threading.Event()
if stop.is_set(): return  # check for cancellation
```

> **Python differs**: no built-in context. Use `asyncio.CancelledError`
> for async code, or `threading.Event` as a manual stop flag for threads.

## Tiny Example

- `main.go` -- WithCancel, WithTimeout, propagation through function calls
- `main.py` -- asyncio task cancellation + threading.Event cancellation

## Common Interview Traps

- **Always call cancel()**: leaked contexts hold resources (always `defer cancel()`)
- **Context is immutable**: `WithCancel` creates a child, doesn't modify the parent
- **Check ctx.Done() in loops**: long-running goroutines must check for cancellation
- **Don't store context in structs**: pass it as the first function parameter
- **WithValue is not a config store**: use it only for request-scoped data (trace IDs)

## What to Say in Interviews

- "I always pass context as the first parameter and defer cancel to prevent leaks"
- "Cancelling a parent context cancels all children -- it propagates down the tree"
- "I check ctx.Done() in select inside long-running loops to respond to cancellation"

## Run It

```bash
go run ./06_concurrency/11_context_cancellation/
python ./06_concurrency/11_context_cancellation/main.py
```

## TL;DR (Interview Summary)

- `ctx, cancel := context.WithCancel(parent)` -- always `defer cancel()`
- `ctx.Done()` returns a channel that closes when cancelled
- `ctx.Err()` returns `Canceled` or `DeadlineExceeded`
- Cancellation **propagates** from parent to all children
- Pass context as first param: `func Foo(ctx context.Context, ...)`
- Don't store context in structs; don't use WithValue as a config store
- Python: `asyncio` task cancellation or `threading.Event`
