# Factory

## What It Is

- **Factory function**: returns an interface based on a configuration value (string, enum, config)
- **ELI10:** A factory is a vending machine for objects -- press "email" and out pops an EmailSender, press "sms" and out pops an SMSSender.
- The caller gets an interface back -- doesn't know or care which concrete type it is

## Why It Matters

- Factories centralize creation logic: add a new sender type without changing callers
- **ELI10:** Without a factory, every caller builds objects themselves -- that's like every restaurant guest cooking their own meal.
- Interviewers check whether you can decouple object creation from usage

## Syntax Cheat Sheet

```go
// Go: factory returns interface
func NewSender(kind string) (Sender, error) {
    switch kind {
    case "email": return &EmailSender{}, nil
    case "sms":   return &SMSSender{}, nil
    default:      return nil, fmt.Errorf("unknown: %s", kind)
    }
}
```

```python
# Python: factory function
def new_sender(kind: str) -> Sender:
    if kind == "email": return EmailSender()
    if kind == "sms":   return SMSSender()
    raise ValueError(f"unknown: {kind}")
```

> Both: factory hides concrete type selection behind a function.

## Tiny Example

- `main.go` -- factory creates Sender from config string; caller uses interface
- `main.py` -- same with dict-based dispatch

## Common Interview Traps

- **Returning concrete type instead of interface**: defeats the purpose of the factory
- **No error case for unknown types**: always handle the default case
- **Using factory when constructor suffices**: don't over-engineer simple creation
- **Fat factories**: if factory does setup + validation + caching, it's doing too much

## What to Say in Interviews

- "My factory returns an interface, so callers don't depend on concrete implementations"
- "Adding a new sender type means adding one case to the factory -- no caller changes"
- "I use factories when the concrete type depends on runtime config like env vars or flags"

## Run It

```bash
go run ./10_design_patterns_in_go/02_factory/
python ./10_design_patterns_in_go/02_factory/main.py
```

## TL;DR (Interview Summary)

- Factory: function that returns an interface based on config
- Caller uses the interface -- doesn't know the concrete type
- Switch/map inside factory; error for unknown types
- Use when concrete type depends on runtime configuration
- Don't use when there's only one implementation (just use a constructor)
- Go: return `(Interface, error)` -- always handle errors
