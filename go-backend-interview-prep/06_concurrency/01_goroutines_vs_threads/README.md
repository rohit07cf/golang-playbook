# Goroutines vs Threads

## What It Is

- A **goroutine** is a lightweight function execution managed by the Go runtime (~2 KB stack)
- Launch with `go f()` -- the runtime multiplexes goroutines onto OS threads

## Why It Matters

- Go can run **millions** of goroutines; OS threads cap around thousands
- Interviewers ask "what's the difference between a goroutine and a thread?" constantly

## Syntax Cheat Sheet

```go
// Go: launch a goroutine
go func() {
    fmt.Println("running concurrently")
}()
```

```python
# Python: threading.Thread (OS thread, GIL limits parallelism)
import threading
t = threading.Thread(target=lambda: print("running concurrently"))
t.start()
```

> **Python differs**: threads are real OS threads but the GIL prevents
> true CPU parallelism. Use `multiprocessing` for CPU-bound work.

## Tiny Example

- `main.go` -- launches several goroutines, prints from each, waits for completion
- `main.py` -- same with `threading.Thread`, shows the GIL limitation note

## Common Interview Traps

- **Goroutine is NOT a thread**: it's a user-space coroutine scheduled by the runtime
- **Stack growth**: goroutines start at ~2 KB and grow dynamically; threads have fixed ~1 MB stacks
- **`go` doesn't guarantee order**: goroutines run concurrently, output order is non-deterministic
- **Main exits = all goroutines die**: if `main()` returns, running goroutines are killed
- **Goroutine leaks**: launching without a way to stop = memory leak over time

## What to Say in Interviews

- "Goroutines are multiplexed onto OS threads by the Go runtime scheduler (M:N scheduling)"
- "They start with a tiny stack (~2 KB) that grows as needed, so millions can coexist"
- "I always ensure goroutines have a termination path -- via channel, context, or WaitGroup"

## Run It

```bash
go run ./06_concurrency/01_goroutines_vs_threads/
python ./06_concurrency/01_goroutines_vs_threads/main.py
```

## TL;DR (Interview Summary)

- `go f()` launches a goroutine -- no return value, no handle
- Goroutine = user-space, ~2 KB; thread = OS-level, ~1 MB
- Go runtime scheduler: M goroutines on N OS threads (M:N model)
- Main exits = all goroutines killed (no daemon concept)
- Always provide a stopping mechanism (channel, context, WaitGroup)
- Python threads are OS threads but GIL limits CPU parallelism
