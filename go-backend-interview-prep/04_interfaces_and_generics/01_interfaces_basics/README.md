# Interfaces Basics

## What It Is

- An interface defines a set of method signatures
- Any type that implements those methods **implicitly** satisfies the interface

## Why It Matters

- Go's polymorphism model -- no `implements` keyword, no inheritance
- "Accept interfaces, return concrete types" is the Go design mantra

## Syntax Cheat Sheet

```go
type Shape interface {
    Area() float64
}

type Circle struct { Radius float64 }

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Circle implicitly satisfies Shape -- no declaration needed
var s Shape = Circle{Radius: 5}
```

**Go vs Python**

```
Go:  type Shape interface { Area() float64 }   // implicit satisfaction
Py:  class Shape(Protocol): def area(self) -> float: ...  # structural typing
```

## What main.go + main.py Show

- Defining an interface and two types that satisfy it
- Passing different concrete types where the interface is expected

## Common Interview Traps

- There is no `implements` keyword -- satisfaction is implicit
- An interface value holds two things: a concrete type and a concrete value
- A nil pointer stored in an interface makes the **interface non-nil**
- You can only call methods defined in the interface, not the concrete type's extras
- Interfaces with pointer-receiver methods require a pointer (method sets)

## What to Say in Interviews

- "Go interfaces are satisfied implicitly, which enables loose coupling."
- "I follow 'accept interfaces, return concrete types' for flexible APIs."
- "Interface values contain a type-value pair, which is why a nil pointer in an interface is not nil."

## Run It

```bash
go run ./04_interfaces_and_generics/01_interfaces_basics
```

```bash
python ./04_interfaces_and_generics/01_interfaces_basics/main.py
```

## TL;DR (Interview Summary)

- Interface = set of method signatures, satisfied implicitly
- No `implements` keyword -- just define the methods
- Interface value = (concrete type, concrete value) pair
- Nil pointer in interface != nil interface
- "Accept interfaces, return concrete types"
- Prefer small interfaces (1-2 methods)
