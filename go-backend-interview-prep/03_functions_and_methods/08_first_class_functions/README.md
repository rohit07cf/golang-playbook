# First-Class Functions

## What It Is

- Functions in Go are **values**: they can be assigned, passed, and returned
- **Function types** define signatures that can be used as parameters

## Why It Matters

- Enables callbacks, middleware, strategy pattern, and functional-style code
- Understanding function types is key to reading Go standard library code

## Syntax Cheat Sheet

```go
// Function type
type Transform func(int) int

// Accept function as parameter
func apply(nums []int, fn Transform) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        result[i] = fn(v)
    }
    return result
}

// Return function
func multiplier(factor int) Transform {
    return func(n int) int { return n * factor }
}
```

**Go vs Python**

```
Go:  type Fn func(int) int           // named function type
Py:  Fn = Callable[[int], int]       # type alias (typing module)
```

## What main.go Shows

- Defining and using function types
- Passing functions as arguments (map/filter pattern)
- Returning functions (factory pattern)
- Simple strategy pattern using function types

## Common Interview Traps

- Function types must match exactly (parameter types + return types)
- You can define named function types for readability
- `nil` is a valid value for a function variable (calling it panics)
- Functions cannot be compared with `==` (except to `nil`)

## What to Say in Interviews

- "Go functions are first-class: they can be stored in variables, passed as args, and returned."
- "I use function types for callbacks, middleware chains, and strategy pattern."
- "Named function types improve readability when the signature is repeated."

## Run It

```bash
go run ./03_functions_and_methods/08_first_class_functions
```

```bash
python ./03_functions_and_methods/08_first_class_functions/main.py
```

## TL;DR (Interview Summary)

- Functions are values: assign, pass, return them
- Named function types: `type Transform func(int) int`
- Use for callbacks, middleware, strategy pattern
- Function types must match signature exactly
- `nil` function value panics when called
- Cannot compare functions with `==` (except to `nil`)
