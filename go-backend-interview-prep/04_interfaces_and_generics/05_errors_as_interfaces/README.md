# Errors as Interfaces

## What It Is

- `error` is a built-in interface with one method: `Error() string`
- **ELI10:** The error interface is just one method: Error() string. That's the whole contract -- even a string can be an error if it tries.
- Any type that has `Error() string` is an error

## Why It Matters

- Go error handling is fundamentally interface-based -- no exceptions
- Custom errors, wrapping, and `errors.Is`/`errors.As` are top interview topics

## Syntax Cheat Sheet

```go
// The error interface
type error interface { Error() string }

// Custom error type
type NotFoundError struct { ID string }
func (e *NotFoundError) Error() string {
    return "not found: " + e.ID
}

// Wrapping
return fmt.Errorf("fetch failed: %w", err)

// Checking
errors.Is(err, ErrNotFound)
errors.As(err, &target)
```

**Go vs Python**

```
Go:  type MyError struct{}; func (e *MyError) Error() string { ... }
Py:  class MyError(Exception): ...
```

## What main.go + main.py Show

- Implementing the `error` interface with a custom type
- Wrapping errors with `%w` and unwrapping with `errors.Is` / `errors.As`
- Sentinel errors

## Common Interview Traps

- `error` is just an interface -- any type with `Error() string` qualifies
- `errors.Is` unwraps the chain; `==` does not
- `errors.As` finds a specific error type in the chain
- `fmt.Errorf("...%w", err)` wraps; `%v` does not (loses the chain)
- A nil `*MyError` stored in an `error` interface is NOT nil
- Sentinel errors (like `io.EOF`) are compared with `errors.Is`, not `==`

## What to Say in Interviews

- "error is a one-method interface; I create custom types to carry context."
- "I wrap errors with %w so callers can use errors.Is and errors.As."
- "I use sentinel errors for expected conditions like EOF, and custom types for structured errors."

## Run It

```bash
go run ./04_interfaces_and_generics/05_errors_as_interfaces
```

```bash
python ./04_interfaces_and_generics/05_errors_as_interfaces/main.py
```

## TL;DR (Interview Summary)

- `error` = interface with `Error() string`
- Custom errors: implement the interface on a struct
- Wrap with `fmt.Errorf("...: %w", err)` -- `%w` preserves the chain
- `errors.Is(err, target)` checks the chain (not just top-level)
- `errors.As(err, &target)` finds a specific type in the chain
- Sentinel errors for expected conditions (like `io.EOF`)
- Nil pointer in error interface is NOT nil
