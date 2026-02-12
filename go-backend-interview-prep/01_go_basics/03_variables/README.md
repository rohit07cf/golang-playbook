# Variables

## What It Is

- Two ways to declare variables: `var` and `:=` (short declaration)
- Every variable in Go has a **zero value** if not explicitly initialized

## Why It Matters

- Zero values prevent uninitialized-variable bugs found in other languages
- Knowing when to use `var` vs `:=` is a basic interview signal

## Syntax Cheat Sheet

```go
// Long form with explicit type
var name string = "alice"

// Type inferred from value
var count = 10

// Zero value (no initializer)
var total int           // total == 0

// Short declaration (inside functions only)
score := 100

// Multiple declarations
var x, y int = 1, 2
a, b := "hello", true
```

## What main.go Shows

- All declaration styles side by side
- Zero values for each type
- The shadowing trap with `:=` in inner scopes

## Common Interview Traps

- `:=` only works inside functions, not at package level
- Unused variables are a **compile error** -- Go enforces this
- `:=` with multiple vars can accidentally shadow an outer variable
- `var x int` gives zero value `0`, not `nil` (nil is for pointers, maps, etc.)
- Re-declaring with `:=` in a new scope creates a new variable (shadowing)

## What to Say in Interviews

- "Go gives every type a zero value -- int is 0, string is empty, bool is false."
- "I use := inside functions for brevity, var at package level for clarity."
- "Unused variables are compile errors, which keeps code clean."

## Run It

```bash
go run ./01_go_basics/03_variables
```

## TL;DR (Interview Summary)

- `var x int` -- explicit type, gets zero value
- `x := 10` -- short declaration, type inferred, functions only
- Zero values: `0` (int), `""` (string), `false` (bool), `nil` (pointers)
- Unused variables = compile error
- `:=` can shadow variables in inner scopes -- watch out
- Use `var` at package level, `:=` inside functions
