# Maps

## What It Is

- Go's built-in hash map: unordered key-value pairs
- Keys must be comparable types (no slices, no maps, no functions as keys)

## Why It Matters

- Maps are the primary lookup structure in Go
- The nil-map-write panic and iteration-order randomness are top interview traps

## Syntax Cheat Sheet

```go
// Literal
m := map[string]int{"alice": 1, "bob": 2}

// make
m2 := make(map[string]int)
m2 := make(map[string]int, 100)  // hint: expect ~100 entries

// Set
m["charlie"] = 3

// Get (returns zero value if missing)
v := m["missing"]   // v == 0

// Comma-ok idiom (check existence)
v, ok := m["alice"]

// Delete
delete(m, "bob")

// Length
len(m)
```

## What main.go Shows

- Creating, reading, writing, deleting map entries
- The comma-ok idiom to distinguish "missing" from "zero value"
- Iteration order is random

## Common Interview Traps

- **Nil map reads** return the zero value (no panic) -- but **nil map writes panic**
- Iteration order is **deliberately randomized** by the runtime
- Maps are **not safe for concurrent access** -- use `sync.Mutex` or `sync.Map`
- You cannot take the address of a map value: `&m["key"]` does not compile
- Keys must be comparable: slices, maps, and functions cannot be keys
- Map lookup always returns a value -- use comma-ok to check existence

## What to Say in Interviews

- "I always use the comma-ok idiom to distinguish a missing key from a zero value."
- "Maps are not goroutine-safe; I protect them with a mutex or use sync.Map."
- "Map iteration order is randomized, so I never depend on it."

## Run It

```bash
go run ./02_data_structures/04_maps
```

## TL;DR (Interview Summary)

- `map[K]V` -- unordered hash map, keys must be comparable
- Nil map read = zero value (safe); nil map write = panic
- Always use comma-ok: `v, ok := m[key]`
- Iteration order is random -- never rely on it
- Not goroutine-safe -- protect with `sync.Mutex`
- `delete(m, key)` is safe even if key is missing
- Cannot take address of map values
