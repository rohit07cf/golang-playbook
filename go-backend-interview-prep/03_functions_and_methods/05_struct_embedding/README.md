# Struct Embedding

## What It Is

- Go uses **embedding** instead of inheritance: one struct includes another
- **ELI10:** Embedding is like getting your parent's toolkit -- you didn't build the tools, but you can use them as if they're yours.
- The embedded struct's methods are **promoted** to the outer struct

## Why It Matters

- This is how Go does composition over inheritance
- **ELI10:** Go skipped inheritance on purpose -- embedding gives you the useful parts without the family drama.
- Interviewers test whether you understand promoted methods and field shadowing

## Syntax Cheat Sheet

```go
type Base struct { ID int }
func (b Base) Describe() string { return fmt.Sprintf("ID=%d", b.ID) }

type User struct {
    Base           // embedded (not a named field)
    Name string
}

u := User{Base: Base{ID: 1}, Name: "Alice"}
u.Describe()   // promoted method -- called on User directly
u.ID           // promoted field -- accessed directly
```

**Go vs Python**

```
Go:  type User struct { Base; Name string }   // embedding (composition)
Py:  class User(Base): ...                     # inheritance
```

## What main.go Shows

- Embedding a struct and accessing promoted methods/fields
- Field shadowing when outer struct has same-name field
- Embedding multiple structs

## Common Interview Traps

- Embedding is **composition**, not inheritance -- there is no "is-a" relationship
- If outer struct declares a field with the same name, it **shadows** the embedded one
- Ambiguity error if two embedded structs export the same method name
- Embedded struct's methods receive the embedded struct as receiver, not the outer
- You can still access shadowed fields via the embedded type name: `u.Base.ID`

## What to Say in Interviews

- "Go uses embedding for composition; there is no inheritance."
- "Promoted methods belong to the embedded type, not the outer type."
- "If I need to override behavior, I define the same method on the outer type."

## Run It

```bash
go run ./03_functions_and_methods/05_struct_embedding
```

```bash
python ./03_functions_and_methods/05_struct_embedding/main.py
```

## TL;DR (Interview Summary)

- Embedding = composition, not inheritance
- Embedded fields and methods are promoted to the outer struct
- Shadowing: outer field with same name hides embedded field
- Access shadowed via explicit name: `u.Base.Field`
- Embedded method receiver is the inner type, not the outer
- Ambiguity if two embeddings export the same name
