# Data Structures -- Quick Revision

Last-minute interview sheet. One screen. No fluff.

---

## Arrays vs Slices

- Array: fixed size, value type, size is part of the type
- Slice: dynamic, reference type (pointer + len + cap), built on an array

```go
arr := [3]int{1, 2, 3}    // array (fixed)
slc := []int{1, 2, 3}     // slice (dynamic)
```

## Slice Internals

```go
s := make([]int, 3, 8)    // len=3, cap=8
s = append(s, 42)          // ALWAYS reassign
sub := s[1:3]              // shares backing array!
```

- When `len == cap`, `append` allocates a new backing array
- Sub-slices share memory -- use `copy` for independence

## Map Gotchas

```go
m := map[string]int{"a": 1}
v, ok := m["missing"]       // v=0, ok=false (comma-ok)
delete(m, "a")               // safe even if key missing
// var nilMap map[string]int
// nilMap["x"] = 1           // PANIC: nil map write
```

- Iteration order is random
- Not goroutine-safe -- use `sync.Mutex`

## Strings vs Runes

```go
s := "Go言語"
len(s)          // 8 (bytes)
len([]rune(s))  // 4 (characters)
for i, r := range s { }  // i=byte index, r=rune
```

- `s[i]` = byte, not character
- Strings are immutable -- convert to `[]rune` to modify

## Pointers

```go
x := 42
p := &x        // *int
*p = 100       // x is now 100
// var np *int  // nil -- dereferencing panics
```

- No pointer arithmetic
- Struct fields auto-deref: `p.Name` (not `(*p).Name`)

## JSON Tags

```go
type User struct {
    Name string `json:"name"`
    Pass string `json:"-"`              // excluded
    Age  int    `json:"age,omitempty"`  // skip if zero
}
```

- Only exported fields are marshaled
- `json.Unmarshal` needs a pointer: `&target`

---

## TL;DR

- Slices = pointer + len + cap; always reassign `append`
- Sub-slices share backing array; use `copy` for independence
- Nil map read = zero value (safe); nil map write = panic
- `len(string)` = bytes, `len([]rune(string))` = characters
- `&` = address-of, `*` = deref; nil deref = panic
- JSON: `omitempty` skips zero values, `-` excludes entirely
- Only exported (uppercase) struct fields appear in JSON
