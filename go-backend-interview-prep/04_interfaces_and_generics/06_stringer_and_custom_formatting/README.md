# Stringer and Custom Formatting

## What It Is

- `fmt.Stringer` is an interface: `String() string`
- **ELI10:** Stringer is like teaching your struct to introduce itself -- "Hi, I'm User{name: Alice}" instead of a memory dump.
- Any type implementing it gets custom output from `fmt.Println`, `%v`, `%s`

## Why It Matters

- Controls how your types display in logs, debug output, and tests
- Shows understanding of implicit interface satisfaction in practice

## Syntax Cheat Sheet

```go
type fmt.Stringer interface {
    String() string
}

type User struct { Name string; Age int }

func (u User) String() string {
    return fmt.Sprintf("%s (age %d)", u.Name, u.Age)
}

fmt.Println(User{Name: "Alice", Age: 30})
// Output: Alice (age 30)
```

**Go vs Python**

```
Go:  func (u User) String() string { ... }   // fmt.Stringer
Py:  def __str__(self) -> str: ...            # print() uses this
```

## What main.go + main.py Show

- Implementing `fmt.Stringer` to customize print output
- How `fmt.Println` automatically calls `String()`
- GoStringer for `%#v` output

## Common Interview Traps

- If `String()` has a **pointer receiver**, only `*T` triggers it in fmt functions
- Avoid infinite recursion: do NOT call `fmt.Sprintf("%v", u)` inside `String()` on the same type
- `GoString()` controls `%#v` output (used for debug)
- `fmt.Stringer` is not the only formatting interface -- there is also `fmt.Formatter`
- Implementing `String()` affects `%v` and `%s` format verbs

## What to Say in Interviews

- "I implement fmt.Stringer to control how my types display in logs and debug output."
- "I'm careful about receiver type -- pointer receiver means only *T triggers String()."
- "I avoid infinite recursion by not using %v on the same type inside String()."

## Run It

```bash
go run ./04_interfaces_and_generics/06_stringer_and_custom_formatting
```

```bash
python ./04_interfaces_and_generics/06_stringer_and_custom_formatting/main.py
```

## TL;DR (Interview Summary)

- `fmt.Stringer`: implement `String() string` for custom output
- `fmt.Println` and `%v`/`%s` automatically call `String()`
- Pointer receiver: only `*T` triggers it in fmt
- Avoid infinite recursion in `String()` -- do not use `%v` on self
- `GoString()` controls `%#v` (debug representation)
- Python equivalent: `__str__` and `__repr__`
