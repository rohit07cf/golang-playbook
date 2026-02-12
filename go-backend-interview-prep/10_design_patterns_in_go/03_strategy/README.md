# Strategy Pattern

## What It Is

- **Strategy**: define a family of algorithms (interfaces), encapsulate each one, and swap them at runtime
- In Go: an interface with one method; each strategy is a struct implementing that interface

## Why It Matters

- Eliminates long if/else or switch chains for choosing behavior
- Interviewers test whether you can decouple algorithm selection from usage

## Syntax Cheat Sheet

```go
// Go: strategy interface
type PricingStrategy interface {
    Calculate(basePrice float64) float64
}
type DiscountPricing struct{ Percent float64 }
func (d DiscountPricing) Calculate(base float64) float64 {
    return base * (1 - d.Percent/100)
}
```

```python
# Python: strategy via callable or class
class DiscountPricing:
    def __init__(self, percent: float):
        self.percent = percent
    def calculate(self, base: float) -> float:
        return base * (1 - self.percent / 100)
```

> Go uses interfaces; Python can use classes or plain functions as strategies.

## Tiny Example

- `main.go` -- NotificationService with swappable routing strategy (priority, round-robin)
- `main.py` -- same with duck-typed strategy objects

## Common Interview Traps

- **Hardcoding the strategy**: the whole point is runtime swapping -- don't hardcode
- **Huge strategy interface**: keep it to 1-2 methods; large interfaces reduce flexibility
- **Strategy vs Factory confusion**: factory creates objects; strategy swaps behavior
- **Not passing context**: strategies may need context (user, config) -- pass it as params

## What to Say in Interviews

- "I use the strategy pattern to swap algorithms at runtime without modifying the caller"
- "Each strategy implements a single-method interface; the service holds the interface"
- "This eliminates switch/if chains and makes adding new strategies trivial"

## Run It

```bash
go run ./10_design_patterns_in_go/03_strategy/
python ./10_design_patterns_in_go/03_strategy/main.py
```

## TL;DR (Interview Summary)

- Strategy = interface with one method; multiple implementations
- Service holds the interface, not the concrete type
- Swap at runtime: `svc.strategy = newStrategy`
- Eliminates if/else chains for algorithm selection
- Keep strategy interface small (1-2 methods)
- Strategy swaps behavior; Factory creates objects -- different patterns
