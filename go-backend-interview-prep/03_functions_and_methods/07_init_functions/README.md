# Init Functions

## What It Is

- `func init()` runs automatically when a package is loaded
- **ELI10:** init() is the alarm clock that rings before main() wakes up -- it runs automatically, once, no questions asked.
- No explicit call needed; runs before `main()`

## Why It Matters

- Used for package-level setup (registering drivers, validating config)
- **ELI10:** If your package needs the lights on before anyone walks in, init() flips the switch.
- Interviewers test whether you know the execution order and pitfalls

## Syntax Cheat Sheet

```go
package main

func init() {
    // runs automatically before main()
    // no parameters, no return values
}

func main() {
    // runs after all init() functions
}
```

**Go vs Python**

```
Go:  func init() { ... }              // auto-runs at package load
Py:  # module-level code runs at import time (similar)
```

## What main.go Shows

- Multiple init functions in the same file (valid in Go)
- Execution order: package-level vars, then init(), then main()

## Common Interview Traps

- A file can have **multiple** `init()` functions (they run in order of appearance)
- `init()` takes no args and returns nothing
- Cannot call `init()` explicitly
- Package-level variables are initialized before `init()` runs
- Import order determines `init()` order across packages
- Overuse of `init()` makes testing hard (side effects at import time)

## What to Say in Interviews

- "init runs automatically at package load, before main, with no parameters."
- "I use init sparingly -- mainly for registering drivers or validating config."
- "For testability, I prefer explicit initialization functions over init."

## Run It

```bash
go run ./03_functions_and_methods/07_init_functions
```

```bash
python ./03_functions_and_methods/07_init_functions/main.py
```

## TL;DR (Interview Summary)

- `func init()` runs automatically before `main()`
- Multiple init functions per file are allowed (run in order)
- No params, no return value, cannot be called explicitly
- Package vars initialized first, then init, then main
- Import order determines cross-package init order
- Use sparingly -- prefer explicit setup for testability
