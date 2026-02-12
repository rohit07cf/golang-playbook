# Go Basics in 5 Minutes

Last-minute revision sheet. One screen. No fluff.

---

## Types

```go
int  float64  bool  string  byte(uint8)  rune(int32)
```

- No implicit conversions: `float64(myInt)` required
- `int` size is platform-dependent

## Variables

```go
var x int = 10      // explicit type
var y = 20          // inferred
z := 30             // short (functions only)
```

- Zero values: `0`, `""`, `false`, `nil`
- Unused variables = compile error

## Constants

```go
const pi = 3.14
const (
    A = iota  // 0
    B         // 1
    C         // 2
)
```

- `iota` resets per `const` block
- Cannot take address of constants

## If / Switch / For

```go
// if with init
if val := compute(); val > 0 { ... }

// switch (no fallthrough by default)
switch x {
case 1, 2: ...
default:   ...
}

// for: the ONLY loop
for i := 0; i < n; i++ { }   // C-style
for n > 0 { }                 // while-style
for { break }                 // infinite
```

- No ternary operator
- No `while` keyword

## Functions

```go
// multiple returns
func f() (int, error) { return 1, nil }

// variadic
func sum(nums ...int) int { ... }

// call with slice spread
sum(mySlice...)
```

- All params passed by value
- No overloading, no default params
- Error is always the last return value

## Defer / Panic / Recover

```go
defer cleanup()          // runs at function exit, LIFO

panic("crash")           // unrecoverable error

defer func() {           // recover must be in defer
    if r := recover(); r != nil { ... }
}()
```

- Defer args evaluated immediately
- Prefer error returns over panic

## Packages

- Uppercase = exported (`Printf`)
- Lowercase = unexported (`helper`)
- Unused imports = compile error
- No circular imports

---

## TL;DR

- Types: int, float64, bool, string, byte, rune -- no implicit casts
- `:=` inside functions, `var` at package level
- `iota` for enums, resets per const block
- `for` is the only loop; `switch` has no fallthrough
- Functions return `(value, error)` -- always check err
- `defer` = cleanup at exit (LIFO); `panic` = crash; `recover` = catch
- Uppercase = public, lowercase = private, unused = compile error
