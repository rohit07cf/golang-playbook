# Wrap Errors and Unwrap

## What It Is

- **Wrapping**: `fmt.Errorf("context: %w", err)` adds context while preserving the original error
- **ELI10:** Wrapping errors is like adding sticky notes to a package as it passes through departments -- each layer adds context.
- **Unwrapping**: `errors.Is` and `errors.As` walk the error chain to find wrapped causes

## Why It Matters

- Real code has call stacks -- wrapping adds context at each layer
- Interviewers test whether you know `%w` vs `%v` and how `errors.Is` / `errors.As` work

## Syntax Cheat Sheet

```go
// Go: wrap with %w, check with Is/As
wrapped := fmt.Errorf("open config: %w", err)
errors.Is(wrapped, ErrNotFound)       // checks chain
errors.As(wrapped, &target)           // extracts typed error
```

```python
# Python: raise ... from preserves cause
try:
    open("missing.txt")
except FileNotFoundError as e:
    raise RuntimeError("open config") from e
# Access: e.__cause__
```

## Tiny Example

- `main.go` -- wraps errors through 3 layers, then uses `errors.Is` and `errors.As` to inspect
- `main.py` -- chains exceptions with `raise ... from`, inspects `__cause__`

## Common Interview Traps

- **Using `%v` instead of `%w`**: `%v` creates a new string, `%w` preserves the chain
- **Wrapping nil errors**: always check `err != nil` before wrapping
- **Double wrapping**: wrapping an already-wrapped error is fine, the chain grows
- **`errors.Is` vs `==`**: `==` only checks the top level; `errors.Is` walks the chain

## What to Say in Interviews

- "I wrap errors with `%w` to add call-site context while preserving the original for `errors.Is`"
- "`errors.Is` checks the chain for a sentinel; `errors.As` extracts a typed error from the chain"
- "In Python the equivalent is `raise ... from` which sets `__cause__`"

## Run It

```bash
go run ./05_errors_and_testing/02_wrap_errors_and_unwrap/
python ./05_errors_and_testing/02_wrap_errors_and_unwrap/main.py
```

## TL;DR (Interview Summary)

- Wrap with `fmt.Errorf("context: %w", err)` -- always use `%w`
- `errors.Is(err, target)` walks the chain for sentinel matches
- `errors.As(err, &target)` walks the chain and extracts a typed error
- `errors.Unwrap(err)` returns the next error in the chain (rarely used directly)
- Python: `raise X from Y` sets `__cause__`; check with `isinstance`
- Never wrap a nil error
