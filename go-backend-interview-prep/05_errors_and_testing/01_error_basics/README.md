# Error Basics

## What It Is

- Go's `error` is a built-in **interface**: `type error interface { Error() string }`
- **ELI10:** Errors in Go are polite "something went wrong" notes. Panic is flipping the table.
- Functions return errors as the **last return value**; callers check `if err != nil`

## Why It Matters

- Every Go function that can fail returns an error -- no hidden exceptions
- **ELI10:** Go makes you handle errors like a responsible adult -- no sweeping exceptions under the rug.
- Interviewers expect you to explain why Go chose explicit error returns over try/catch

## Syntax Cheat Sheet

```go
// Go: return error as last value
result, err := strconv.Atoi("abc")
if err != nil {
    log.Fatal(err)
}
```

```python
# Python: raise/except for the same purpose
try:
    result = int("abc")
except ValueError as e:
    print(e)
```

## Tiny Example

- `main.go` -- divides two numbers, returns an error for division by zero, shows `if err != nil`
- `main.py` -- same logic using `raise ValueError` and `try/except`

## Common Interview Traps

- **Ignoring the error**: `val, _ := someFunc()` silently drops failures
- **Checking err after using the value**: always check `err != nil` first
- **Returning nil error with non-zero value**: if err is nil, the value must be valid
- **Confusing nil interface vs nil pointer**: a non-nil interface can hold a nil pointer (typed nil trap)

## What to Say in Interviews

- "Go uses explicit error returns instead of exceptions -- every call site decides how to handle failure"
- "The `error` interface has one method: `Error() string` -- any type can be an error"
- "I always check `err != nil` before using the success value"

## Run It

```bash
go run ./05_errors_and_testing/01_error_basics/
python ./05_errors_and_testing/01_error_basics/main.py
```

## TL;DR (Interview Summary)

- `error` is an interface: `{ Error() string }`
- Functions return `(result, error)` -- always check err first
- Never ignore errors with `_` in production code
- Go chose values over exceptions for **explicit, local** error handling
- Python uses `try/except`; Go uses `if err != nil`
