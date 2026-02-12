# Observer / Events Pattern

## What It Is

- **Observer**: objects subscribe to events; when an event fires, all subscribers are notified
- **ELI10:** Observer is a newsletter subscription -- you sign up once and get notified whenever something happens.
- In Go: an event bus holds a list of handler functions; `Publish()` calls them all

## Why It Matters

- Decouples the publisher from subscribers: adding a new reaction requires zero publisher changes
- **ELI10:** Without observer, the sender must know everyone who cares -- with observer, it just shouts into a megaphone and listeners tune in.
- Interviewers test event-driven thinking for microservices and async workflows

## Syntax Cheat Sheet

```go
// Go: simple event bus
type EventBus struct {
    handlers map[string][]func(Event)
}
func (b *EventBus) Subscribe(event string, fn func(Event)) { ... }
func (b *EventBus) Publish(event string, data Event) { ... }
```

```python
# Python: event bus with dict of lists
class EventBus:
    def __init__(self):
        self.handlers: dict[str, list] = {}
    def subscribe(self, event, fn): ...
    def publish(self, event, data): ...
```

> Both: publish/subscribe pattern decouples event producers from consumers.

## Tiny Example

- `main.go` -- event bus for notification events; logger + analytics subscribers
- `main.py` -- same with callable subscribers

## Common Interview Traps

- **Synchronous subscribers blocking**: long handlers delay the publisher
- **No error handling in subscribers**: one panic/exception shouldn't crash the bus
- **Memory leaks**: subscribers that never unsubscribe hold references
- **Order dependency**: don't rely on subscriber execution order

## What to Say in Interviews

- "I use an event bus to decouple components: the sender publishes, subscribers react independently"
- "Adding a new subscriber requires zero changes to the publisher"
- "For production, I'd make the bus async with goroutines or a message queue"

## Run It

```bash
go run ./10_design_patterns_in_go/07_observer_events/
python ./10_design_patterns_in_go/07_observer_events/main.py
```

## TL;DR (Interview Summary)

- Event bus: `Subscribe(event, handler)` + `Publish(event, data)`
- Publisher doesn't know who subscribes -- fully decoupled
- Adding new behavior = adding a subscriber, zero publisher changes
- Handle errors in subscribers: one failure shouldn't crash the bus
- For production: async via goroutines or message queues
- Don't rely on subscriber execution order
