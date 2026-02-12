# Functions (Intro)

## What It Is

- Named blocks of code with typed parameters and return values
- The basic unit of behavior in Go

## Why It Matters

- Everything in Go revolves around functions -- even methods are functions with a receiver
- Knowing function signatures is essential for reading any Go code

## Syntax Cheat Sheet

```go
// Basic function
func greet(name string) string {
    return "hello, " + name
}

// Multiple parameters of the same type (shorthand)
func add(a, b int) int {
    return a + b
}

// No return value
func log(msg string) {
    fmt.Println(msg)
}

// Functions are first-class (can be assigned to variables)
fn := func(x int) int { return x * 2 }
```

**Go vs Python**
Go:  `func add(a, b int) int { return a + b }`
Py:  `def add(a: int, b: int) -> int: return a + b`

## What main.go Shows

- Defining and calling basic functions
- Parameter shorthand for same-type params
- Assigning a function to a variable

## Common Interview Traps

- Parameters are passed **by value** (copies) -- not by reference
- Slices, maps, and channels look like pass-by-reference but are actually reference types passed by value
- Go does not support default parameter values
- Go does not support function overloading (same name, different args)
- Unused function parameters do NOT cause compile errors (only unused variables do)

## What to Say in Interviews

- "Go passes everything by value, but slices and maps contain internal pointers, so mutations are visible to the caller."
- "There is no function overloading in Go -- each function has a unique name."
- "Functions are first-class in Go: they can be assigned, passed, and returned."

## Run It

```bash
go run ./01_go_basics/08_functions_intro
```

```bash
python ./01_go_basics/08_functions_intro/main.py
```

## TL;DR (Interview Summary)

- Functions: `func name(params) returnType { ... }`
- Same-type params shorthand: `func add(a, b int)`
- All parameters are passed by value
- No default parameters, no overloading
- Functions are first-class values
- Even methods are just functions with a receiver
