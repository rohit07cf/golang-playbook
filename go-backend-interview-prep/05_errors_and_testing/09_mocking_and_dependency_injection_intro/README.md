# Mocking and Dependency Injection Intro

## What It Is

- **Dependency injection**: pass dependencies (interfaces) into functions/structs instead of hard-coding them
- **Mocking**: create a fake implementation of an interface for testing
- **ELI10:** Mocking is hiring a stunt double -- looks like the real thing, but it's safe and predictable. DI is like bringing your own charger to a hotel -- the room doesn't care what brand, as long as it fits the outlet.

## Why It Matters

- Real code depends on databases, APIs, filesystems -- you can't call those in unit tests
- Interviewers test whether you know how to design testable code with interfaces

## Syntax Cheat Sheet

```go
// Go: define interface, inject it
type UserStore interface {
    GetUser(id int) (string, error)
}
type Service struct { store UserStore }

// Test: pass a fake
type FakeStore struct{}
func (f FakeStore) GetUser(id int) (string, error) { return "alice", nil }
```

```python
# Python: protocol or duck typing, pass a stub
class FakeStore:
    def get_user(self, user_id: int) -> str:
        return "alice"

service = UserService(store=FakeStore())
```

## Tiny Example

- `main.go` -- defines `UserStore` interface, `UserService` struct, and a `FakeStore` for testing
- `main.py` -- same pattern using a protocol-style class and a stub

## Common Interview Traps

- **Mocking without interfaces**: if you depend on a concrete type, you can't swap it in tests
- **Too many mocks**: if everything is mocked, you're testing wiring, not logic
- **Mock returning only happy path**: always test error cases too
- **Huge interfaces**: small, focused interfaces (1-3 methods) are easier to mock

## What to Say in Interviews

- "I inject dependencies as interfaces so I can swap in fakes during testing"
- "Small interfaces are easier to mock -- I follow the interface segregation principle"
- "In Go, no `implements` keyword is needed -- the fake just needs to have the right methods"

## Run It

```bash
go run ./05_errors_and_testing/09_mocking_and_dependency_injection_intro/
python ./05_errors_and_testing/09_mocking_and_dependency_injection_intro/main.py
```

## TL;DR (Interview Summary)

- Define small interfaces for external dependencies
- Inject them as constructor/function parameters
- Create fake/stub implementations for tests
- Go: interface is satisfied implicitly (no `implements`)
- Test both happy path and error path with fakes
- Python: duck typing or Protocol -- same idea, less ceremony
