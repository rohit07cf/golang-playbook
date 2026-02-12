# Builder vs Functional Options

## What It Is

- **Builder**: step-by-step object construction with chained methods
- **Functional options**: `WithXxx()` functions that modify a config -- the Go-idiomatic approach

## Why It Matters

- Functional options is the #1 Go-specific design pattern interviewers ask about
- It solves the "constructor with 10 optional params" problem cleanly

## Syntax Cheat Sheet

```go
// Go: functional options
type Option func(*NotificationService)
func WithTimeout(d time.Duration) Option {
    return func(s *NotificationService) { s.timeout = d }
}
func NewService(opts ...Option) *NotificationService {
    s := &NotificationService{timeout: 5 * time.Second}
    for _, opt := range opts { opt(s) }
    return s
}
// Usage: NewService(WithTimeout(10*time.Second), WithRetries(3))
```

```python
# Python: **kwargs with defaults (natural equivalent)
class NotificationService:
    def __init__(self, timeout=5, retries=1, **kwargs):
        self.timeout = timeout
        self.retries = retries
```

> **Go**: functional options for optional config. **Python**: `**kwargs` with defaults.

## Tiny Example

- `main.go` -- functional options for NotificationService; shows builder alternative for comparison
- `main.py` -- `**kwargs` defaults as the natural Python equivalent

## Common Interview Traps

- **Using builder pattern in Go**: functional options is more idiomatic
- **Non-obvious defaults**: always document what the zero-value/default is
- **Too many options**: if you have 15+ options, the struct may be doing too much
- **Mutating after construction**: options are applied during construction only

## What to Say in Interviews

- "I use functional options for configurable constructors: `NewService(WithTimeout(10s))`"
- "Each option is a closure that modifies the struct -- composable and self-documenting"
- "This avoids the builder pattern's chaining and gives clear defaults"

## Run It

```bash
go run ./10_design_patterns_in_go/09_builder_vs_functional_options/
python ./10_design_patterns_in_go/09_builder_vs_functional_options/main.py
```

## TL;DR (Interview Summary)

- Functional options: `WithXxx()` closures applied in constructor
- `func NewService(opts ...Option)` -- variadic, composable, clear defaults
- Each option modifies the struct: `func(s *Service) { s.field = val }`
- Builder works but is less idiomatic in Go
- Python: `**kwargs` with defaults solves the same problem natively
- Use when constructor has 3+ optional parameters
- Don't use for required parameters -- those stay as regular constructor args
