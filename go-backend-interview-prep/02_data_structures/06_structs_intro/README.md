# Structs (Intro)

## What It Is

- A struct is a composite type that groups named fields together
- Go's version of classes -- but with no inheritance

## Why It Matters

- Structs are how you model data in Go (users, requests, configs, etc.)
- Understanding value semantics and zero values is essential

## Syntax Cheat Sheet

```go
// Define a struct type
type User struct {
    Name  string
    Email string
    Age   int
}

// Create instances
u1 := User{Name: "alice", Email: "a@b.com", Age: 30}
u2 := User{}                // zero value: {"", "", 0}
u3 := User{Name: "bob"}    // partial: Email="", Age=0

// Access fields
fmt.Println(u1.Name)

// Anonymous struct (one-off use)
point := struct{ X, Y int }{10, 20}
```

**Go vs Python**
Go:  type User struct { Name string }
Py:  @dataclass class User: name: str = ""

## What main.go Shows

- Defining struct types and creating instances
- Zero-value structs and partial initialization
- Structs are value types (assigning copies all fields)
- Anonymous structs for one-off use

## Common Interview Traps

- Structs are **value types** -- assigning copies the entire struct
- Zero-value struct has each field set to its type's zero value
- Field names starting with uppercase are exported; lowercase are unexported
- You cannot compare structs with `==` if they contain non-comparable fields (slices, maps)
- No constructors -- use `NewXxx()` factory functions by convention
- Embedding is not inheritance (covered in later modules)

## What to Say in Interviews

- "Structs in Go are value types; assigning or passing them makes a full copy."
- "I use NewXxx factory functions to enforce invariants since Go has no constructors."
- "Exported fields start with an uppercase letter -- this is how Go controls visibility."

## Run It

```bash
go run ./02_data_structures/06_structs_intro
```

```bash
python ./02_data_structures/06_structs_intro/main.py
```

## TL;DR (Interview Summary)

- `type Name struct { ... }` defines a composite type
- Structs are value types -- assigning copies everything
- Zero value: all fields get their type's zero value
- Uppercase field = exported; lowercase = unexported
- No constructors -- use `NewXxx()` factory functions
- Comparable with `==` only if all fields are comparable
- Use pointers (`*User`) to avoid copying large structs
