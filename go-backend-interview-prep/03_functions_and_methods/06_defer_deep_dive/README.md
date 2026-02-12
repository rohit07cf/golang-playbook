# Defer Deep Dive

## What It Is

- `defer` schedules a function call to run when the enclosing function returns
- This is a deeper look: loop traps, argument evaluation, and production patterns

## Why It Matters

- Defer-in-loop is a classic production bug (resource leak)
- Understanding argument evaluation timing is a common interview question

## Syntax Cheat Sheet

```go
// Basic defer
defer f.Close()

// Defer with closure (captures current value)
defer func() { fmt.Println(x) }()

// Defer evaluates args NOW, not at exit
x := 10
defer fmt.Println(x)  // prints 10, even if x changes later
```

**Go vs Python**

```
Go:  defer f.Close()                 // runs at function exit
Py:  with open(path) as f: ...       # context manager closes on exit
```

## What main.go Shows

- Argument evaluation timing
- Defer in loops (the trap and the fix)
- Using defer with closures to capture current state

## Common Interview Traps

- Defer args evaluated **at the defer statement**, not at function exit
- Defer in a loop: resources pile up and are only released when the function returns
- Fix: extract the loop body into a helper function so defer runs per iteration
- Deferred functions run even if the function panics
- LIFO order: last deferred = first to run
- Named return values can be modified by a deferred closure

## What to Say in Interviews

- "Defer arguments are evaluated immediately, but the call executes at function exit."
- "I avoid defer in loops because resources accumulate; I use a helper function instead."
- "Deferred closures can modify named return values, which is useful for error enrichment."

## Run It

```bash
go run ./03_functions_and_methods/06_defer_deep_dive
```

```bash
python ./03_functions_and_methods/06_defer_deep_dive/main.py
```

## TL;DR (Interview Summary)

- Defer args evaluated at defer statement, not at exit
- LIFO execution order
- Defer in loops = resource leak; extract to helper function
- Deferred closures can modify named return values
- Deferred functions run even during panic
- Use defer for cleanup: file close, mutex unlock, connection release
