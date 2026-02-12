# Hexagonal Architecture (Ports & Adapters) Mini

## What It Is

- **Ports**: interfaces that define how the core talks to the outside world (sender, repository)
- **Adapters**: concrete implementations of those interfaces (email adapter, postgres adapter)

## Why It Matters

- Core business logic depends only on ports (interfaces) -- fully testable, swappable
- Interviewers ask about clean architecture to test your separation-of-concerns thinking

## Syntax Cheat Sheet

```go
// Go: port = interface (defined in core)
type SenderPort interface { Send(to, msg string) error }
type RepoPort interface   { Save(n Notification) error }
// Adapter = implementation (in adapter package)
type EmailAdapter struct{ ... }
func (e *EmailAdapter) Send(to, msg string) error { ... }
```

```python
# Python: port = duck type / Protocol
class SenderPort(Protocol):
    def send(self, to: str, msg: str) -> None: ...
# Adapter = class implementing the protocol
class EmailAdapter:
    def send(self, to, msg): ...
```

> Core logic imports zero external packages. Adapters live outside core.

## Tiny Example

- `main.go` -- core service with SenderPort + RepoPort; email/memory adapters; all wired in main
- `main.py` -- same layering with duck-typed ports

## Common Interview Traps

- **Core importing adapters**: core depends on ports only -- adapters depend on core, not vice versa
- **Too many layers**: hexagonal doesn't mean 10 packages -- keep it practical
- **Confusing ports with adapters**: port = interface; adapter = implementation
- **Not testing the core independently**: the whole point is core testability

## What to Say in Interviews

- "I structure my app with ports and adapters: the core defines interfaces, adapters implement them"
- "The core has zero external dependencies -- I can test it with in-memory fakes"
- "main() wires adapters to ports, so the dependency direction always points inward"

## Run It

```bash
go run ./10_design_patterns_in_go/11_hexagonal_ports_adapters_mini/
python ./10_design_patterns_in_go/11_hexagonal_ports_adapters_mini/main.py
```

## TL;DR (Interview Summary)

- Port = interface defined by core logic
- Adapter = concrete implementation of a port
- Core depends on ports only -- zero external imports
- Adapters depend on core (implement its interfaces)
- main() wires adapters to ports (dependency injection)
- Testable: swap adapters for fakes without touching core
- Don't over-layer: 3 layers (core, ports, adapters) is enough
