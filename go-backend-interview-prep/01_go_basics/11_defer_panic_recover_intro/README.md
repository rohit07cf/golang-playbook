# Defer, Panic, Recover (Intro)

## What It Is

- **defer**: schedules a function call to run when the enclosing function returns
- **panic**: crashes the program (like an unhandled exception)
- **recover**: catches a panic, preventing the crash

## Why It Matters

- `defer` is used everywhere for cleanup (closing files, unlocking mutexes)
- Knowing when to use panic vs error return is a key interview signal

## Syntax Cheat Sheet

```go
// Defer: runs at function exit, LIFO order
func readFile() {
    f, _ := os.Open("data.txt")
    defer f.Close()   // guaranteed cleanup
    // ... use f ...
}

// Panic: crash the program
panic("something went terribly wrong")

// Recover: catch a panic (must be inside a deferred function)
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()
```

## What main.go Shows

- Defer execution order (LIFO)
- A panic being caught by recover
- Why defer is preferred for cleanup

## Common Interview Traps

- Deferred calls execute in **LIFO** order (last deferred = first executed)
- Defer arguments are evaluated **immediately**, not at execution time
- `recover()` only works inside a **deferred** function -- nowhere else
- `panic` should be used for truly unrecoverable situations, not for normal errors
- Deferred functions run even if the function panics

## What to Say in Interviews

- "I use defer for cleanup like closing files and releasing locks."
- "Panic is for unrecoverable bugs, not for expected errors -- those use error returns."
- "Recover only works inside a deferred function, and I rarely use it directly."

## Run It

```bash
go run ./01_go_basics/11_defer_panic_recover_intro
```

## TL;DR (Interview Summary)

- `defer` runs at function exit, LIFO order
- Defer args are evaluated at the defer statement, not at execution
- Use defer for cleanup: file close, mutex unlock, connection release
- `panic` = crash; use only for unrecoverable situations
- `recover()` catches panics; must be in a deferred function
- Prefer returning errors over panicking
