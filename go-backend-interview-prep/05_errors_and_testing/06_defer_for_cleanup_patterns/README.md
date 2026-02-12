# Defer for Cleanup Patterns

## What It Is

- `defer` schedules a function call to run when the enclosing function returns
- **ELI10:** Defer for cleanup is like always putting your tools back in the shed before going home -- even if it rained.
- Used for **cleanup**: closing files, releasing locks, recovering from panics

## Why It Matters

- Resource leaks are a common production bug; defer prevents them
- Interviewers expect you to know defer ordering (LIFO) and the file-close idiom

## Syntax Cheat Sheet

```go
// Go: defer runs at function return (LIFO order)
f, err := os.Open("data.txt")
if err != nil { return err }
defer f.Close()
```

```python
# Python: context manager (with statement) or try/finally
with open("data.txt") as f:
    data = f.read()

# Or explicit try/finally
f = open("data.txt")
try:
    data = f.read()
finally:
    f.close()
```

## Tiny Example

- `main.go` -- file close, mutex unlock, multi-defer LIFO order, recover-in-defer
- `main.py` -- context managers, `try/finally`, same patterns

## Common Interview Traps

- **Defer in a loop**: each iteration adds a new deferred call -- may exhaust resources before the function returns
- **Defer evaluates args immediately**: `defer fmt.Println(x)` captures x at the defer line, not at return
- **Defer after error check**: always check `err` before deferring `Close()`
- **Ignoring Close() error**: `defer f.Close()` discards the error; for writes, check it

## What to Say in Interviews

- "I always defer Close() immediately after a successful Open() to prevent leaks"
- "Deferred calls run in LIFO order, and arguments are evaluated when defer is called"
- "In Python I'd use a context manager (`with`) for the same guarantee"

## Run It

```bash
go run ./05_errors_and_testing/06_defer_for_cleanup_patterns/
python ./05_errors_and_testing/06_defer_for_cleanup_patterns/main.py
```

## TL;DR (Interview Summary)

- `defer f.Close()` right after a successful open -- never before the error check
- Deferred calls execute in **LIFO** order (last deferred runs first)
- Arguments are evaluated **at the defer line**, not at function return
- Don't defer in loops -- use a helper function or close manually
- For writes, capture the Close() error: `defer func() { err = f.Close() }()`
- Python: `with` statement or `try/finally`
