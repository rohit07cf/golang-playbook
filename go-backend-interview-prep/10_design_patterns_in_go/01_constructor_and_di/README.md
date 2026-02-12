# Constructor and Dependency Injection

## What It Is

- **Constructor**: a `NewXxx()` function that builds a struct with its dependencies wired in
- **Dependency Injection (DI)**: passing dependencies as arguments instead of creating them internally

## Why It Matters

- DI makes code testable -- swap a real email sender for a fake in tests
- Interviewers check whether you wire dependencies explicitly or hide them behind globals

## Syntax Cheat Sheet

```go
// Go: constructor injection
type NotificationService struct {
    sender Sender   // interface, not concrete type
    logger Logger
}
func NewNotificationService(s Sender, l Logger) *NotificationService {
    return &NotificationService{sender: s, logger: l}
}
```

```python
# Python: constructor injection
class NotificationService:
    def __init__(self, sender: Sender, logger: Logger):
        self.sender = sender
        self.logger = logger
```

> Both languages: dependencies come in through the constructor, not via global imports.

## Tiny Example

- `main.go` -- NotificationService with injected Sender and Logger; swap implementations
- `main.py` -- same pattern with duck-typed dependencies

## Common Interview Traps

- **Importing concrete types directly**: makes unit testing impossible without mocks
- **Using init() for DI**: init() runs at package load -- no control, hard to test
- **Huge constructors**: if you need 8+ deps, the struct does too much -- split it
- **Not using interfaces for deps**: accept interfaces, return structs
- **Confusing DI with a DI framework**: Go doesn't need Wire/Dig -- just pass arguments

## What to Say in Interviews

- "I inject dependencies through the constructor so I can swap them in tests"
- "My constructors accept interfaces, not concrete types -- this decouples the implementation"
- "Go doesn't need a DI framework; explicit wiring in main() is clear and testable"

## Run It

```bash
go run ./10_design_patterns_in_go/01_constructor_and_di/
python ./10_design_patterns_in_go/01_constructor_and_di/main.py
```

## TL;DR (Interview Summary)

- `NewXxx(deps)` constructor wires dependencies explicitly
- Accept interfaces, return structs
- DI = testable: swap real sender for a fake in tests
- No DI framework needed in Go -- explicit wiring in main()
- If constructor has 8+ params, the struct does too much
- `init()` is not for DI -- avoid hidden initialization
