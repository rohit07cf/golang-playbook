# Slice Tricks

## What It Is

- Common operations on slices that Go does not provide built-in (delete, insert, filter)
- Patterns every Go developer uses regularly

## Why It Matters

- There is no `slice.delete()` in Go -- you must know the idioms
- Interviewers test whether you understand the underlying memory effects

## Syntax Cheat Sheet

```go
// Delete element at index i (order preserved)
s = append(s[:i], s[i+1:]...)

// Delete element at index i (order NOT preserved, faster)
s[i] = s[len(s)-1]
s = s[:len(s)-1]

// Insert at index i
s = append(s[:i], append([]T{val}, s[i:]...)...)

// Filter in-place (no allocation)
n := 0
for _, v := range s {
    if keep(v) {
        s[n] = v
        n++
    }
}
s = s[:n]

// Copy a slice (independent)
dst := make([]int, len(src))
copy(dst, src)
```

**Go vs Python**
Go:  s = append(s[:i], s[i+1:]...)  // delete
Py:  del s[i]                        # built-in

## What main.go Shows

- Deleting an element (preserving and not preserving order)
- Inserting into the middle of a slice
- Filtering in place
- Preallocating with `make` to avoid repeated allocations

## Common Interview Traps

- Deleting from a slice without understanding backing array leaves stale pointers (for pointer slices)
- The "fast delete" (swap with last) changes order -- only use when order does not matter
- `append(s[:i], s[i+1:]...)` modifies the original backing array
- Forgetting to preallocate: `make([]T, 0, n)` avoids O(n) reallocations
- Inserting into a slice is O(n) because elements shift -- there is no O(1) insert
- `slices.Delete` was added in Go 1.21 as a standard library function

## What to Say in Interviews

- "I use append with re-slicing for delete since Go has no built-in remove."
- "For large slices I preallocate with make to avoid repeated allocations during append."
- "I choose between order-preserving and swap-with-last delete depending on requirements."

## Run It

```bash
go run ./02_data_structures/03_slice_tricks
```

```bash
python ./02_data_structures/03_slice_tricks/main.py
```

## TL;DR (Interview Summary)

- Delete (keep order): `s = append(s[:i], s[i+1:]...)`
- Delete (fast, no order): swap with last, then shrink
- Insert at i: re-slice + append
- Filter in-place: overwrite-and-shrink pattern
- Preallocate: `make([]T, 0, n)` to avoid repeated growth
- All slice operations modify the backing array -- copy first if you need the original
