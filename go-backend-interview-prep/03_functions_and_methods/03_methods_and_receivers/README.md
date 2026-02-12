# Methods and Receivers

## What It Is

- A **method** is a function with a special receiver argument
- **Value receiver** `(r Rect)`: operates on a copy
- **Pointer receiver** `(r *Rect)`: operates on the original

## Why It Matters

- Value vs pointer receiver is the #1 methods question in Go interviews
- Getting this wrong causes silent bugs (mutations lost on copies)

## Syntax Cheat Sheet

```go
type Rect struct { Width, Height float64 }

// Value receiver -- cannot mutate
func (r Rect) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver -- can mutate
func (r *Rect) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

**Go vs Python**

```
Go:  func (r Rect) Area() float64 { return r.Width * r.Height }
Py:  def area(self) -> float: return self.width * self.height
```

## What main.go Shows

- Defining methods with value and pointer receivers
- Showing that value receivers cannot mutate, pointer receivers can
- Go auto-dereferences: you can call a pointer-receiver method on a value

## Common Interview Traps

- Value receiver: method gets a **copy** -- mutations are lost
- Pointer receiver: method gets the **original** -- mutations persist
- Go auto-takes address: `val.PointerMethod()` works (compiler adds `&`)
- Convention: if any method needs a pointer receiver, make them all pointer receivers
- Methods can only be defined on types in the same package
- You cannot define methods on built-in types directly (use a named type)

## What to Say in Interviews

- "I use pointer receivers when the method needs to mutate state or the struct is large."
- "Value receivers are for read-only methods on small structs."
- "If any method uses a pointer receiver, I make all methods on that type use pointer receivers for consistency."

## Run It

```bash
go run ./03_functions_and_methods/03_methods_and_receivers
```

```bash
python ./03_functions_and_methods/03_methods_and_receivers/main.py
```

## TL;DR (Interview Summary)

- Value receiver `(r T)` = copy, cannot mutate the original
- Pointer receiver `(r *T)` = reference, can mutate
- Go auto-takes address for pointer receiver methods
- If one method needs `*T`, make all methods use `*T`
- Methods can only be defined in the same package as the type
- Cannot define methods on built-in types -- create a named type
