# Closures

## What It Is

- An anonymous function that captures variables from its enclosing scope
- The captured variables persist as long as the closure exists

## Why It Matters

- Closures are used everywhere: callbacks, goroutines, functional patterns
- The loop-variable capture trap is one of the most common Go interview questions

## Syntax Cheat Sheet

```go
// Basic closure
counter := 0
increment := func() int {
    counter++
    return counter
}

// Closure as return value
func makeAdder(x int) func(int) int {
    return func(y int) int { return x + y }
}
```

**Go vs Python**

```
Go:  fn := func(x int) int { return x * 2 }
Py:  fn = lambda x: x * 2
```

## What main.go Shows

- Capturing and mutating variables from outer scope
- Returning closures from functions
- The loop-variable trap and how to fix it

## Common Interview Traps

- Closures capture variables **by reference**, not by value
- Classic bug: capturing loop variable `i` -- all closures see the final value
- Fix: pass the variable as a parameter to create a local copy
- Go 1.22+ changed `for` loop variable semantics (each iteration gets a new variable)
- Closures can outlive the scope they were created in

## What to Say in Interviews

- "Closures capture variables by reference, so I'm careful with loop variables."
- "In Go pre-1.22, I shadow the loop variable or pass it as a param to avoid the trap."
- "I use closures for callbacks and to build functions that carry state."

## Run It

```bash
go run ./03_functions_and_methods/01_closures
```

```bash
python ./03_functions_and_methods/01_closures/main.py
```

## TL;DR (Interview Summary)

- Closures capture outer variables by reference
- Loop-variable trap: all closures see the same final value
- Fix: `v := v` inside loop or pass as param
- Go 1.22+ fixes loop variable scoping (each iteration = new var)
- Closures can be returned from functions and carry state
- Used heavily with goroutines and callbacks
