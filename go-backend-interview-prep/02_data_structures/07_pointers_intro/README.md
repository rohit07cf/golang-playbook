# Pointers (Intro)

## What It Is

- A pointer holds the **memory address** of a value
- `&` gets the address; `*` dereferences (reads the value at that address)

## Why It Matters

- Pointers let you modify the original value (not a copy) across function calls
- Critical for efficient struct passing and understanding Go's memory model

## Syntax Cheat Sheet

```go
x := 42
p := &x            // p is *int, points to x

fmt.Println(*p)    // 42 (dereference)
*p = 100           // modifies x through the pointer

// Pointer to a struct
u := &User{Name: "alice"}
fmt.Println(u.Name)  // auto-dereferenced (no -> like C)

// new() allocates and returns a pointer
np := new(int)     // *int, points to zero-value int (0)

// nil pointer
var q *int         // nil
```

**Go vs Python**
Go:  p := &x; *p = 100        // explicit pointer
Py:  # no pointers; mutable objects use reference semantics

## What main.go Shows

- Taking addresses and dereferencing
- Modifying values through pointers
- Passing pointers to functions to mutate the original
- Nil pointer behavior

## Common Interview Traps

- Go has **no pointer arithmetic** (unlike C)
- Dereferencing a nil pointer causes a **runtime panic**
- Struct fields are auto-dereferenced: `p.Name` works (no need for `(*p).Name`)
- `new(T)` allocates a zero-value T and returns `*T` -- rarely used (prefer `&T{}`)
- Pointers to local variables are safe -- Go's escape analysis may put them on the heap
- Maps, slices, and channels already contain internal pointers -- no need to pass `*[]int`

## What to Say in Interviews

- "I pass pointers to structs to avoid copying and to allow mutation."
- "Go has no pointer arithmetic, which prevents a whole class of memory bugs."
- "The compiler's escape analysis decides whether a local variable lives on the stack or heap."

## Run It

```bash
go run ./02_data_structures/07_pointers_intro
```

```bash
python ./02_data_structures/07_pointers_intro/main.py
```

## TL;DR (Interview Summary)

- `&x` = address of x; `*p` = value at address p
- Pass `*T` to modify the original, not a copy
- Struct fields auto-deref: `p.Name` (no `->`)
- Nil pointer dereference = runtime panic
- No pointer arithmetic in Go
- `new(T)` returns `*T` with zero value; prefer `&T{}`
- Escape analysis decides stack vs heap -- you do not manage this manually
