# Constants

## What It Is

- Values fixed at compile time, declared with `const`
- `iota` is a built-in counter for creating enumerated constants

## Why It Matters

- Constants prevent accidental mutation of config values
- `iota` is idiomatic Go for enums -- interviewers expect you to know it

## Syntax Cheat Sheet

```go
// Simple constant
const pi = 3.14159

// Typed constant
const maxRetries int = 5

// Constant block with iota
const (
    Sunday    = iota   // 0
    Monday             // 1
    Tuesday            // 2
)

// Iota with expressions
const (
    KB = 1 << (10 * (iota + 1))   // 1024
    MB                              // 1048576
    GB                              // 1073741824
)
```

**Go vs Python**
Go:  `const Pi = 3.14         // compile-time, iota for enums`
Py:  `PI = 3.14               # convention only, use IntEnum for enums`

## What main.go Shows

- Simple constants, typed constants, and `iota` enums
- Using `iota` with bit-shift expressions for sizes

## Common Interview Traps

- Constants must be determinable at compile time -- no function calls
- Untyped constants are more flexible (can be used in expressions with different types)
- `iota` resets to 0 in each new `const` block
- You cannot take the address of a constant (`&myConst` fails)
- There is no `enum` keyword in Go; `iota` is the pattern

## What to Say in Interviews

- "Go uses iota for enumerations instead of a dedicated enum type."
- "Untyped constants in Go adapt to the context where they are used."
- "Constants are compile-time only -- you cannot use runtime values."

## Run It

```bash
go run ./01_go_basics/04_constants
```

```bash
python ./01_go_basics/04_constants/main.py
```

## TL;DR (Interview Summary)

- `const` declares compile-time values
- `iota` starts at 0, increments by 1 per line in a const block
- `iota` resets in each new `const` block
- Untyped constants adapt to their usage context
- No `&` (address-of) on constants
- Go has no `enum` keyword -- use `const` + `iota`
