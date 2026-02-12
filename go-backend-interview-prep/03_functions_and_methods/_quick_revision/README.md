# Functions and Methods -- Quick Revision

Last-minute interview sheet. One screen. No fluff.

---

## Closures

```go
counter := 0
inc := func() int { counter++; return counter }
```

- Captures by reference -- beware loop variable trap
- Fix: shadow with `i := i` or pass as param

## Recursion

- Always define a base case
- Go has no tail-call optimization -- prefer iteration for large n
- Naive Fibonacci is O(2^n) -- use memoization

## Value vs Pointer Receiver

```go
func (r Rect) Area() float64   // value receiver: copy
func (r *Rect) Scale(f float64) // pointer receiver: mutates
```

- Value receiver = read-only
- Pointer receiver = mutation
- If any method needs `*T`, use `*T` for all

## Method Sets

- `T` method set: value-receiver methods only
- `*T` method set: value + pointer receiver methods
- Interface assignment checks the method set

## Struct Embedding

```go
type Admin struct { User; Level string }
```

- Composition, not inheritance
- Promoted methods belong to inner type
- Shadowing: outer field hides inner field with same name

## Defer

```go
defer f.Close()              // runs at function exit, LIFO
defer fmt.Println(x)        // x captured NOW
defer func() { ... }()      // closure sees current values at exit
```

- Defer in loops = resource leak; extract to helper function
- Deferred closures can modify named return values

## Init Functions

- `func init()` runs before `main()`, no params, no return
- Multiple init functions per file allowed
- Order: pkg vars -> init() -> main()

---

## TL;DR

- Closures capture by reference; shadow loop vars to fix
- Value receiver = copy; pointer receiver = mutation
- `*T` method set includes all methods; `T` only value-receiver methods
- Embedding = composition; promoted methods belong to inner type
- Defer args evaluated immediately; defer in loops = trap
- init() runs before main, automatically -- use sparingly
- Functions are first-class values: pass, return, store
