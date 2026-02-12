# Generics vs Interfaces

## What It Is

- Two tools for polymorphism: interfaces (runtime dispatch) vs generics (compile-time)
- Knowing when to use each is the key interview question

## Why It Matters

- Overusing generics makes code complex; overusing interfaces loses type safety
- "When would you use generics vs interfaces?" is a direct interview question

## Syntax Cheat Sheet

```go
// Interface: runtime polymorphism, behavior-based
type Writer interface { Write([]byte) (int, error) }

// Generic: compile-time polymorphism, type-based
func Sort[T cmp.Ordered](s []T) { ... }

// Rule of thumb:
// - Interface: "I need behavior X" (methods)
// - Generic: "I need the same logic for different types" (containers, algorithms)
```

**Go vs Python**

```
Go:  interface = runtime dispatch; generic = compile-time monomorphization
Py:  Protocol/ABC = runtime duck typing; generics = type hints (not enforced)
```

## What main.go + main.py Show

- Same problem solved with interfaces vs generics
- When each approach is cleaner
- The decision framework

## Common Interview Traps

- Generics are NOT always better than interfaces
- Interfaces work with dynamic dispatch (slightly slower, more flexible)
- Generics are resolved at compile time (faster, less flexible)
- You cannot have generic methods (only generic functions and types)
- Interfaces are best for behavior; generics are best for containers/algorithms
- Do not use generics just because they exist -- simplicity wins

## What to Say in Interviews

- "I use interfaces when I need to abstract over behavior, like io.Reader."
- "I use generics when I have identical logic for different types, like a type-safe collection."
- "If an interface works and is simpler, I prefer it. I only add generics when type safety matters and interfaces are insufficient."

## Run It

```bash
go run ./04_interfaces_and_generics/09_generics_vs_interfaces
```

```bash
python ./04_interfaces_and_generics/09_generics_vs_interfaces/main.py
```

## TL;DR (Interview Summary)

- **Interface**: abstract over behavior (methods). Runtime dispatch.
- **Generic**: same logic for different types. Compile-time resolution.
- Use interfaces for: handlers, readers, writers, services
- Use generics for: containers, algorithms, transformations
- Generics cannot have methods (only functions and types)
- Prefer simplicity -- if an interface works, use it
- "When in doubt, start with an interface"
