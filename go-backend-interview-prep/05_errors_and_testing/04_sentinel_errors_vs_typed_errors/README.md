# Sentinel Errors vs Typed Errors

## What It Is

- **Sentinel errors**: package-level `var ErrX = errors.New("...")` -- simple, named values
- **Typed errors**: struct implementing `Error()` -- carries structured data

## Why It Matters

- Choosing the right error strategy is a design decision interviewers probe
- Using the wrong one leads to brittle code or unnecessary complexity

## Syntax Cheat Sheet

```go
// Go sentinel
var ErrNotFound = errors.New("not found")
errors.Is(err, ErrNotFound)

// Go typed
type NotFoundError struct{ ID int }
func (e *NotFoundError) Error() string { ... }
errors.As(err, &target)
```

```python
# Python: exception classes serve both roles
class NotFoundError(Exception):
    pass  # sentinel-like (no fields)

class DetailedNotFoundError(Exception):
    def __init__(self, resource_id: int):
        self.resource_id = resource_id  # typed (has fields)
```

## Tiny Example

- `main.go` -- defines both sentinel and typed errors, shows when to use `errors.Is` vs `errors.As`
- `main.py` -- equivalent using exception classes with and without fields

## Common Interview Traps

- **Exporting sentinels makes them part of your API**: changing them is a breaking change
- **Comparing typed errors with `==`**: use `errors.As`, not `==`
- **Creating sentinels when you need data**: if you need context fields, use a typed error
- **Overusing typed errors**: for "yes/no" failures (EOF, not found), sentinels are simpler
- **Forgetting `errors.Is` walks the chain**: `==` only checks the top level

## What to Say in Interviews

- "I use sentinel errors for well-known failure modes like `io.EOF` or `ErrNotFound`"
- "I use typed errors when callers need structured data like status codes or field names"
- "Sentinels use `errors.Is`, typed errors use `errors.As` -- both walk wrapped chains"

## Run It

```bash
go run ./05_errors_and_testing/04_sentinel_errors_vs_typed_errors/
python ./05_errors_and_testing/04_sentinel_errors_vs_typed_errors/main.py
```

## TL;DR (Interview Summary)

- **Sentinel**: `var ErrX = errors.New(...)` -- check with `errors.Is`
- **Typed**: struct + `Error()` -- check with `errors.As`
- Use sentinels for simple yes/no failure modes
- Use typed errors when callers need structured data
- Exporting a sentinel makes it part of your public API
- Python uses exception class hierarchy for both patterns
