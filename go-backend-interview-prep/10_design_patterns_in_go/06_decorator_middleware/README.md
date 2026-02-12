# Decorator / Middleware Pattern

## What It Is

- **Decorator**: wraps an object to add behavior (logging, retry, rate limiting) without changing the original
- In Go: a function that takes an interface and returns the same interface with added behavior

## Why It Matters

- Every Go HTTP middleware uses this pattern: `func(http.Handler) http.Handler`
- Interviewers expect you to layer decorators without modifying the core implementation

## Syntax Cheat Sheet

```go
// Go: decorator wraps Sender interface
func WithLogging(next Sender, logger Logger) Sender {
    return &loggingSender{next: next, logger: logger}
}
// loggingSender implements Sender, delegates to next
```

```python
# Python: decorator wraps sender
class LoggingSender:
    def __init__(self, next_sender, logger):
        self.next = next_sender
        self.logger = logger
    def send(self, to, msg):
        self.logger.log(f"sending to {to}")
        self.next.send(to, msg)
```

> Both: decorator has the same interface as the wrapped object. Callers don't know it's wrapped.

## Tiny Example

- `main.go` -- layer logging + retry decorators around a sender
- `main.py` -- same with class-based wrappers

## Common Interview Traps

- **Decorator changes the interface**: it must keep the same interface as the wrapped object
- **Confusing with adapter**: adapter changes interface shape; decorator adds behavior
- **Order matters**: logging(retry(sender)) logs the retry; retry(logging(sender)) does not
- **Too many layers**: 5+ decorators become hard to debug -- use sparingly

## What to Say in Interviews

- "I wrap the sender with decorators like logging and retry -- each keeps the Sender interface"
- "Decorators compose: `WithLogging(WithRetry(sender))` -- order determines behavior"
- "This is the same pattern as Go HTTP middleware: `func(Handler) Handler`"

## Run It

```bash
go run ./10_design_patterns_in_go/06_decorator_middleware/
python ./10_design_patterns_in_go/06_decorator_middleware/main.py
```

## TL;DR (Interview Summary)

- Decorator wraps an interface, adds behavior, delegates to the original
- Same interface in and out -- callers don't know it's wrapped
- Compose: `WithLogging(WithRetry(baseSender))`
- Order matters: outermost decorator runs first
- HTTP middleware is this exact pattern: `func(http.Handler) http.Handler`
- Adapter changes shape; decorator adds behavior -- different patterns
