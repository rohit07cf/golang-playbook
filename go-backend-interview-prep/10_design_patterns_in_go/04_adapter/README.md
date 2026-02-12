# Adapter Pattern

## What It Is

- **Adapter**: wraps a third-party or legacy API to match your internal interface
- The adapter translates between two incompatible interfaces

## Why It Matters

- Third-party APIs change -- an adapter isolates the blast radius to one file
- Interviewers test whether you can decouple your code from external dependencies

## Syntax Cheat Sheet

```go
// Go: adapter wraps external API behind your interface
type Sender interface {
    Send(to, msg string) error
}
type TwilioAdapter struct{ client *twilio.Client }
func (t *TwilioAdapter) Send(to, msg string) error {
    return t.client.Messages.Create(to, msg)  // translate
}
```

```python
# Python: adapter class
class TwilioAdapter:
    def __init__(self, client):
        self.client = client
    def send(self, to: str, msg: str) -> None:
        self.client.messages.create(to=to, body=msg)
```

> Both: your code calls `Send()`; the adapter translates to the third-party API.

## Tiny Example

- `main.go` -- adapts a "third-party" Twilio-like API to the internal Sender interface
- `main.py` -- same with a simulated external API class

## Common Interview Traps

- **Adapting too much**: adapter should only translate; no business logic inside
- **Leaking third-party types**: never expose third-party types in your interface
- **Not testing the adapter**: test with a fake external client
- **Confusing adapter with decorator**: adapter changes the interface; decorator adds behavior

## What to Say in Interviews

- "I wrap third-party APIs behind my own interface so my code doesn't depend on them directly"
- "If the external API changes, I only update the adapter -- no changes to business logic"
- "Adapter changes the interface shape; decorator keeps the same interface and adds behavior"

## Run It

```bash
go run ./10_design_patterns_in_go/04_adapter/
python ./10_design_patterns_in_go/04_adapter/main.py
```

## TL;DR (Interview Summary)

- Adapter wraps a third-party API behind your own interface
- Your code depends on your interface, not the external one
- Adapter translates method signatures -- no business logic
- External API changes? Update one adapter file
- Adapter != decorator: adapter changes shape; decorator adds behavior
- Always hide third-party types behind your own types
