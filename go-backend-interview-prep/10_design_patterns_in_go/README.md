# 10 -- Design Patterns in Go

## What This Module Covers

Go doesn't have classes or inheritance, so classic OOP patterns look very different.
Go-idiomatic design relies on small interfaces, composition, and explicit dependency
injection rather than abstract base classes and singletons.

This module teaches the patterns that actually appear in production Go codebases --
from constructor injection and functional options to repository layers and middleware
chains. Every example uses a consistent "Notification Service" domain to keep
cognitive load low.

Every example includes `main.go` + Python equivalent `main.py`.

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `constructor_and_di` | Constructor injection, testable dependencies |
| 02 | `factory` | Factory functions that return interface implementations |
| 03 | `strategy` | Swap algorithms at runtime via interfaces |
| 04 | `adapter` | Wrap third-party code behind your own interface |
| 05 | `repository` | Data access interface + in-memory fake for tests |
| 06 | `decorator_middleware` | Wrap behavior: logging, retry, rate limiting |
| 07 | `observer_events` | Publish/subscribe event bus |
| 08 | `singleton_and_why_avoid` | sync.Once vs passing dependencies (prefer DI) |
| 09 | `builder_vs_functional_options` | Functional options (Go-idiomatic) vs builder |
| 10 | `pipeline_and_handlers_chain` | Chain of responsibility: validate -> enrich -> send |
| 11 | `hexagonal_ports_adapters_mini` | Ports (interfaces) + adapters (implementations) |

---

## 10-Min Revision Path

1. Read `01_constructor_and_di` -- the foundation: inject, don't import
2. Skim `02_factory` -- factory returns an interface, caller doesn't know the concrete type
3. Glance at `03_strategy` -- swap behavior without if/else chains
4. Read `06_decorator_middleware` -- the pattern behind every Go HTTP middleware
5. Skim `05_repository` -- data access layer that's testable
6. Read `09_builder_vs_functional_options` -- functional options is the Go interview favorite
7. Glance at `11_hexagonal_ports_adapters_mini` -- architecture-level pattern
8. Finish with `_quick_revision/README.md` -- one-screen cheat sheet

---

## Common Design Pattern Mistakes

- Using singletons when constructor injection works (makes testing hard)
- Creating huge interfaces instead of small, focused ones
- Using inheritance thinking: embedding is NOT inheritance
- Overusing factory functions when a simple constructor suffices
- Putting business logic in handlers instead of a service layer
- Not defining interfaces where they're consumed (accept interfaces, return structs)
- Building complex builders when functional options are cleaner
- Skipping the repository pattern and putting SQL in handlers
- Tight coupling to third-party APIs: always wrap behind an interface (adapter)
- Over-engineering: not every service needs every pattern

---

## TL;DR

- **DI > singleton**: inject dependencies via constructors; avoid global state
- **Small interfaces**: 1-2 methods; defined by the consumer, not the implementor
- **Composition > inheritance**: embed structs, don't build deep hierarchies
- **Functional options**: `WithTimeout(5*time.Second)` -- idiomatic Go configuration
- **Repository pattern**: interface for data access; swap implementations for testing
- **Adapter**: wrap external APIs behind your own interface
- **Middleware/decorator**: `func(Handler) Handler` -- composable behavior wrapping
