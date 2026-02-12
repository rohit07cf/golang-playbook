# Design Patterns in Go -- Quick Revision

> One-screen refresher. Skim before your interview.

---

## When to Use Each Pattern (1-liner)

| Pattern | Use When |
|---------|----------|
| Constructor/DI | Always -- inject deps via `NewXxx(deps)` |
| Factory | Concrete type depends on runtime config |
| Strategy | Need to swap algorithms at runtime |
| Adapter | Wrapping a third-party API behind your interface |
| Repository | Abstracting data access for testability |
| Decorator | Adding behavior (logging, retry) without changing the original |
| Observer | Multiple components react to events independently |
| Singleton | Almost never -- prefer DI; use `sync.Once` only for truly global resources |
| Functional Options | Constructor with 3+ optional parameters |
| Pipeline | Sequential processing steps: validate -> enrich -> send |
| Hexagonal | Architecture-level: core depends on ports (interfaces), adapters live outside |

## Go-Idiomatic Takeaways

- **Small interfaces**: 1-2 methods; defined where consumed
- **Composition over inheritance**: embed, don't extend
- **Accept interfaces, return structs**
- **Functional options** > builder for optional config
- **DI via constructor** > singleton > global var

## Key Distinctions

- **DI vs Singleton**: DI = explicit, testable; Singleton = hidden, global
- **Repository vs Adapter**: Repo = data access; Adapter = API translation
- **Strategy vs Factory**: Strategy swaps behavior; Factory creates objects
- **Decorator vs Adapter**: Decorator adds behavior (same interface); Adapter changes shape

## 11 Tiny Go Snippets

```go
// 1. Constructor DI
func NewService(s Sender, l Logger) *Service {
    return &Service{sender: s, logger: l}
}
```

```go
// 2. Factory
func NewSender(kind string) (Sender, error) {
    switch kind {
    case "email": return &EmailSender{}, nil
    case "sms":   return &SMSSender{}, nil
    default:      return nil, fmt.Errorf("unknown: %s", kind)
    }
}
```

```go
// 3. Strategy
type Router interface { Route(msg string) string }
svc.SetRouter(RoundRobin{})  // swap at runtime
```

```go
// 4. Adapter
type TwilioAdapter struct{ client *TwilioClient }
func (a *TwilioAdapter) Send(to, msg string) error {
    return a.client.CreateMessage(to, msg, a.from)
}
```

```go
// 5. Repository
type Repo interface {
    Save(n Notification) error
    FindByID(id string) (Notification, error)
}
```

```go
// 6. Decorator
func WithLogging(next Sender) Sender {
    return &loggingSender{next: next}
}
// Chain: WithLogging(WithRetry(base))
```

```go
// 7. Observer
bus.Subscribe("order.created", func(e Event) { log(e) })
bus.Publish(Event{Name: "order.created", Data: data})
```

```go
// 8. Singleton (avoid)
var once sync.Once
func GetConfig() *Config {
    once.Do(func() { cfg = loadConfig() })
    return cfg
}
// Better: cfg := loadConfig(); NewService(cfg)
```

```go
// 9. Functional options
func WithTimeout(d time.Duration) Option {
    return func(s *Service) { s.timeout = d }
}
svc := NewService(sender, WithTimeout(10*time.Second))
```

```go
// 10. Pipeline
pipeline := NewPipeline(
    &ValidateHandler{}, &EnrichHandler{}, &SendHandler{},
)
pipeline.Run(msg) // short-circuits on error
```

```go
// 11. Hexagonal ports
type SenderPort interface { Send(to, msg string) error }
type RepoPort interface   { Save(n Notification) error }
// Core uses ports; main() wires adapters
```

## 11 Tiny Python Snippets

```python
# 1. Constructor DI
class Service:
    def __init__(self, sender, logger):
        self.sender = sender
```

```python
# 2. Factory
def new_sender(kind):
    return {"email": EmailSender, "sms": SMSSender}[kind]()
```

```python
# 3. Strategy
svc.strategy = RoundRobin()  # swap at runtime
```

```python
# 4. Adapter
class TwilioAdapter:
    def send(self, to, msg):
        self.client.create_message(to, msg)
```

```python
# 5. Repository
class MemoryRepo:
    def save(self, n): self._data[n.id] = n
    def find_by_id(self, id): return self._data[id]
```

```python
# 6. Decorator
class LoggingSender:
    def __init__(self, next): self.next = next
    def send(self, to, msg):
        print(f"log: {to}")
        self.next.send(to, msg)
```

```python
# 7. Observer
bus.subscribe("order.created", lambda e: print(e))
bus.publish(Event("order.created", data))
```

```python
# 8. Singleton (avoid)
_config = None
def get_config():
    global _config
    if not _config: _config = load_config()
    return _config
# Better: config = load_config(); Service(config)
```

```python
# 9. Functional options = kwargs
Service(sender, timeout=10, retries=3)
```

```python
# 10. Pipeline
pipeline = Pipeline(ValidateHandler(), EnrichHandler(), SendHandler())
pipeline.run(msg)
```

```python
# 11. Hexagonal
class Service:
    def __init__(self, sender, repo):  # ports = duck-typed params
        self.sender = sender
```

## 11 Interview One-Liners

| # | Pattern | One-Liner |
|---|---------|-----------|
| 1 | DI | "Inject deps via constructor; accept interfaces, return structs" |
| 2 | Factory | "Returns interface based on config; caller doesn't know concrete type" |
| 3 | Strategy | "Swap algorithms at runtime via a single-method interface" |
| 4 | Adapter | "Wraps third-party API behind my interface; isolates external changes" |
| 5 | Repository | "Interface for data access; swap in-memory for SQL without touching services" |
| 6 | Decorator | "Same interface in and out; layers behavior like middleware" |
| 7 | Observer | "Publish/subscribe decouples producers from consumers" |
| 8 | Singleton | "Avoid it; prefer DI. If needed, use sync.Once for thread safety" |
| 9 | Options | "Functional options: `NewService(WithTimeout(10s))` -- idiomatic Go" |
| 10 | Pipeline | "Chain of handlers: validate -> enrich -> send; short-circuit on error" |
| 11 | Hexagonal | "Core depends on ports (interfaces); adapters live outside; main() wires" |

## TL;DR

- **DI always**: inject via constructor, test with fakes
- **Small interfaces**: 1-2 methods, defined by consumer
- **Composition > inheritance**: Go has no classes -- embed and compose
- **Functional options**: the Go answer to optional constructor params
- **Adapter for external APIs**: isolate third-party behind your interface
- **Decorator for behavior**: logging, retry, rate limit -- same interface
- **Repository for data**: interface in service, implementation in adapter
- **Avoid singletons**: hidden deps, testing hell -- use DI instead
