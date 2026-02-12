# Multiple Return Values

## What It Is

- Go functions can return two or more values
- **ELI10:** Multiple returns are like a doctor visit -- you get the diagnosis AND the bill at the same time, no separate trips.
- The most common pattern: `(result, error)`

## Why It Matters

- This is how Go handles errors -- no exceptions, no try/catch
- **ELI10:** Instead of throwing a tantrum (exception), Go hands you a calm note saying "here is what went wrong" right next to the answer.
- Interviewers expect you to know the `(value, error)` pattern cold

## Syntax Cheat Sheet

```go
// Two return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Calling: always check the error
result, err := divide(10, 3)
if err != nil {
    log.Fatal(err)
}

// Ignore a value with _
result, _ = divide(10, 2)
```

**Go vs Python**
Go:  `val, err := divide(10, 3)`
Py:  `val, err = divide(10, 3)      # tuple unpacking`

## What main.go Shows

- Returning two values from a function
- The `(value, error)` pattern
- Using `_` to discard a return value

## Common Interview Traps

- You **must** use or discard every return value -- unused = compile error
- `_` (blank identifier) discards a value intentionally
- Ignoring errors with `_` is valid syntax but bad practice in production
- You cannot return a single value from a function declared to return two
- The error is conventionally the **last** return value

## What to Say in Interviews

- "Go uses multiple return values instead of exceptions for error handling."
- "The convention is (result, error) with error always last."
- "I always check err != nil before using the result."

## Run It

```bash
go run ./01_go_basics/09_multiple_return_values
```

```bash
python ./01_go_basics/09_multiple_return_values/main.py
```

## TL;DR (Interview Summary)

- Functions can return multiple values: `func f() (int, error)`
- Error is always the last return value by convention
- Always check `err != nil` before using the result
- Use `_` to explicitly discard a return value
- This replaces exceptions -- Go has no try/catch
- Ignoring errors silently is a code smell
