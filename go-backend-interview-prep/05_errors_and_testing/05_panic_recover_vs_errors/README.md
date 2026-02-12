# Panic, Recover vs Errors

## What It Is

- **panic**: immediately stops normal execution, unwinds the stack, runs deferred functions
- **ELI10:** Panic is the fire alarm. Recover is the fire extinguisher. Errors are the polite "please exit through the side door."
- **recover**: called inside a deferred function to catch a panic and resume normally
- **errors**: the normal way to signal expected failures

## Why It Matters

- Interviewers test whether you know **when** to panic vs return an error
- Misusing panic for expected failures is a red flag in interviews

## Syntax Cheat Sheet

```go
// Go: panic for bugs, error for expected failures
panic("impossible state")       // programmer bug
recover()                       // catch panic in defer

func riskyOp() error {          // expected failure
    return errors.New("timeout")
}
```

```python
# Python: similar split
raise RuntimeError("impossible state")  # like panic
# except ... catches it (like recover)

def risky_op() -> None:
    raise TimeoutError("timeout")  # expected failure
```

## Tiny Example

- `main.go` -- shows panic for programmer bugs, recover in defer, and normal error returns
- `main.py` -- `raise` + `try/except` for both crash-level and expected failures

## Common Interview Traps

- **Using panic for user input errors**: panic is for bugs, not validation
- **ELI10:** Using panic for normal errors is like calling 911 because you ran out of milk.
- **Recovering in the wrong function**: `recover()` only works in a **deferred** function
- **Panics cross goroutine boundaries**: an unrecovered panic in a goroutine crashes the program
- **Returning after recover**: the function returns the zero value unless you set it in defer
- **Panicking in libraries**: libraries should return errors, not panic (callers can't predict it)

## What to Say in Interviews

- "I use errors for expected failures and panic only for unrecoverable programmer bugs"
- "recover() must be called inside a deferred function; it returns nil if there's no panic"
- "Libraries should never panic -- they should return errors so callers decide what to do"

## Run It

```bash
go run ./05_errors_and_testing/05_panic_recover_vs_errors/
python ./05_errors_and_testing/05_panic_recover_vs_errors/main.py
```

## TL;DR (Interview Summary)

- **error**: for expected failures (file not found, timeout, bad input)
- **panic**: for programmer bugs (index out of range, nil dereference, impossible state)
- `recover()` catches panics but only inside a `defer` function
- An unrecovered panic kills the goroutine (and the program if it's main)
- Libraries should return errors, not panic
- Python: all failures use exceptions, but the idiom split is similar (ValueError vs AssertionError)
