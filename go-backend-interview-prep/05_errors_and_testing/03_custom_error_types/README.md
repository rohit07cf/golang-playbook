# Custom Error Types

## What It Is

- Any struct that implements `Error() string` satisfies the `error` interface
- **ELI10:** A custom error is like a doctor's diagnosis instead of "I feel bad" -- it carries specific, actionable information.
- Custom errors carry **structured data** (status codes, field names, context)

## Why It Matters

- Real applications need errors with metadata, not just strings
- Interviewers ask how to create and check custom errors with `errors.As`

## Syntax Cheat Sheet

```go
// Go: struct + Error() method
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

```python
# Python: custom exception class
class ValidationError(Exception):
    def __init__(self, field: str, message: str):
        self.field = field
        self.message = message
        super().__init__(f"{field}: {message}")
```

## Tiny Example

- `main.go` -- defines `ValidationError` and `HTTPError`, returns and checks them with `errors.As`
- `main.py` -- same pattern using Python custom exception classes and `isinstance`

## Common Interview Traps

- **Pointer vs value receiver**: use `*T` for `Error()` so `errors.As` works with `*T`
- **Returning `*MyError(nil)` as `error`**: the interface is non-nil even if the pointer is nil (typed nil trap)
- **Forgetting `errors.As`**: plain type assertion won't walk a wrapped chain
- **Too many custom types**: only create them when you need structured data; use sentinels for simple cases

## What to Say in Interviews

- "I create custom error types when I need structured fields like status codes or field names"
- "I always use a pointer receiver for `Error()` and check with `errors.As` to walk wrapped chains"
- "For simple known-failure cases, a sentinel error is enough"

## Run It

```bash
go run ./05_errors_and_testing/03_custom_error_types/
python ./05_errors_and_testing/03_custom_error_types/main.py
```

## TL;DR (Interview Summary)

- Implement `Error() string` on a struct to make it an error
- Use pointer receiver (`*T`) -- makes `errors.As` work correctly
- Carry structured data: status codes, field names, retry info
- Check with `errors.As(err, &target)` -- walks the chain
- Python: custom exception classes with attributes
- Don't over-create -- sentinel errors work for simple cases
