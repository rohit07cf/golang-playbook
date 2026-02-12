# Generics Basics

## What It Is

- Generics let you write functions and types that work with **any type** while keeping type safety
- Added in Go 1.18 -- uses type parameters in square brackets `[T any]`

## Why It Matters

- Eliminates duplicate code for different types (before generics: copy-paste or `any`)
- Interviewers test whether you know the syntax and when to reach for generics

## Syntax Cheat Sheet

```go
// Generic function
func Map[T any, U any](s []T, fn func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s { result[i] = fn(v) }
    return result
}

// Generic type
type Stack[T any] struct { items []T }
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
```

**Go vs Python**

```
Go:  func Map[T any, U any](s []T, fn func(T) U) []U
Py:  def map_fn(s: list[T], fn: Callable[[T], U]) -> list[U]: ...
```

## What main.go + main.py Show

- A generic `Map` function that transforms a slice
- A generic `Stack` type with Push/Pop
- Using generics with different concrete types

## Common Interview Traps

- Type params go in square brackets `[T any]`, not angle brackets
- You cannot use `==` on a type parameter unless constrained with `comparable`
- Generic functions are monomorphized at compile time
- You cannot use type assertions on type parameters
- Methods on generic types must repeat the type parameter: `func (s *Stack[T])`

## What to Say in Interviews

- "Generics eliminate duplicate code while preserving type safety at compile time."
- "I reach for generics when I have the same logic for different types -- like containers or transformations."
- "Type parameters use square brackets, and I constrain them to express what operations are needed."

## Run It

```bash
go run ./04_interfaces_and_generics/07_generics_basics
```

```bash
python ./04_interfaces_and_generics/07_generics_basics/main.py
```

## TL;DR (Interview Summary)

- `func Name[T constraint](params)` -- square brackets for type params
- `any` is the loosest constraint (accepts everything)
- Generic types: `type Stack[T any] struct { ... }`
- Methods repeat the param: `func (s *Stack[T]) Push(v T)`
- Resolved at compile time -- no runtime overhead
- Cannot use `==` without `comparable` constraint
- Cannot use type assertions on type parameters
