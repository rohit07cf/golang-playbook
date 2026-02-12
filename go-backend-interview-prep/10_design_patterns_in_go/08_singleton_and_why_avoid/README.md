# Singleton and Why to Avoid It

## What It Is

- **Singleton**: ensures only one instance of a type exists (Go: `sync.Once` + package-level var)
- **ELI10:** A singleton is that one remote everyone fights over -- convenient until two people need it at the same time.
- **Why avoid**: global state makes testing hard, hides dependencies, creates tight coupling

## Why It Matters

- Interviewers ask "when would you use a singleton?" -- the best answer is "rarely"
- **ELI10:** Singletons are like global variables in disguise -- everyone depends on them, and nobody can test in peace.
- Knowing `sync.Once` is useful, but knowing when NOT to use singletons is more important

## Syntax Cheat Sheet

```go
// Go: singleton via sync.Once (thread-safe)
var instance *Config
var once sync.Once
func GetConfig() *Config {
    once.Do(func() { instance = loadConfig() })
    return instance
}
```

```python
# Python: module-level variable (natural singleton)
_config = None
def get_config():
    global _config
    if _config is None:
        _config = load_config()
    return _config
```

> Both: singleton = global state. Prefer passing dependencies explicitly.

## Tiny Example

- `main.go` -- shows singleton pitfall (hard to test), then the DI alternative
- `main.py` -- same comparison: module global vs injected dependency

## Common Interview Traps

- **"Singleton is a design pattern"**: yes, but it's an anti-pattern in most Go code
- **Hidden dependency**: calling `GetInstance()` inside a function hides what it depends on
- **Testing nightmare**: can't replace the singleton in tests without global mutation
- **Concurrency issues**: forgetting `sync.Once` leads to race conditions
- **init() as singleton**: `init()` runs once but makes testing even harder

## What to Say in Interviews

- "I avoid singletons because they hide dependencies and make testing difficult"
- "If I need exactly one instance, I create it in main() and inject it everywhere"
- "sync.Once is useful for lazy init of truly global resources like config, but I prefer DI"

## Run It

```bash
go run ./10_design_patterns_in_go/08_singleton_and_why_avoid/
python ./10_design_patterns_in_go/08_singleton_and_why_avoid/main.py
```

## TL;DR (Interview Summary)

- `sync.Once` makes singleton thread-safe in Go
- Singleton = global state = hidden dependency = testing hell
- Prefer: create in main(), inject via constructor
- Acceptable singletons: config, logger (maybe) -- still prefer DI
- `init()` is worse than sync.Once -- no control over timing
- "Accept interfaces, return structs" eliminates the need for singletons
