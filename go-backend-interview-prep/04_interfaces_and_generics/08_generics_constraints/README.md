# Generics Constraints

## What It Is

- A **constraint** limits which types a type parameter accepts
- **ELI10:** Constraints are the velvet rope at the club -- "You must be at least this comparable to enter this function."
- Built-in: `any`, `comparable`. Custom: interface with type lists

## Why It Matters

- Constraints express what operations you need on a type parameter
- Knowing `comparable`, union constraints, and `~` (underlying type) is interview-relevant

## Syntax Cheat Sheet

```go
// comparable: allows == and !=
func Index[T comparable](s []T, target T) int { ... }

// Custom constraint with type union
type Number interface {
    ~int | ~float64
}
func Sum[T Number](s []T) T { ... }

// ~ means "underlying type" (includes named types based on int)
type MyInt int   // ~int matches MyInt
```

**Go vs Python**

```
Go:  type Number interface { ~int | ~float64 }   // type union constraint
Py:  Number = TypeVar("Number", int, float)       # TypeVar with bounds
```

## What main.go + main.py Show

- Using `comparable` for equality checks
- Defining a custom `Number` constraint
- The `~` operator for underlying types

## Common Interview Traps

- `comparable` allows `==` and `!=` only -- not `<`, `>`
- For ordering, use `cmp.Ordered` (Go 1.21+) or define your own constraint
- `~int` matches any type whose underlying type is `int` (including named types)
- Type unions in constraints are not interfaces you can use as regular types
- Constraints can embed other constraints (interface embedding)

## What to Say in Interviews

- "I use comparable when I need equality, and cmp.Ordered or custom constraints for ordering."
- "The tilde ~ matches underlying types, so named types like type UserID int also qualify."
- "Constraints let me express exactly what operations the generic code needs."

## Run It

```bash
go run ./04_interfaces_and_generics/08_generics_constraints
```

```bash
python ./04_interfaces_and_generics/08_generics_constraints/main.py
```

## TL;DR (Interview Summary)

- `comparable`: allows `==` and `!=`
- `cmp.Ordered`: allows `<`, `>`, `<=`, `>=` (Go 1.21+)
- Custom constraint: `type Number interface { ~int | ~float64 }`
- `~T` matches types with underlying type T
- Type unions work only in constraints, not as regular types
- Constraints can embed other constraints
