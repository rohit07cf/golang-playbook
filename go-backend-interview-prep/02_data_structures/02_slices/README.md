# Slices

## What It Is

- A dynamically-sized, flexible view into an underlying array
- The most commonly used data structure in Go

## Why It Matters

- Slices are everywhere -- function params, return values, iteration
- Understanding **length vs capacity** and the **backing array** is critical for interviews

## Syntax Cheat Sheet

```go
// Literal
nums := []int{1, 2, 3}

// make: length, capacity
s := make([]int, 5)        // len=5, cap=5, zero-filled
s2 := make([]int, 0, 10)   // len=0, cap=10

// Append (always reassign!)
s = append(s, 42)

// Slice a slice (shares backing array)
sub := nums[1:3]           // [2, 3]

// Length and capacity
len(s)
cap(s)

// Nil slice vs empty slice
var nilSlice []int          // nil, len=0, cap=0
emptySlice := []int{}       // not nil, len=0, cap=0
```

**Go vs Python**
Go:  s := make([]int, 0, 10)  // len=0, cap=10
Py:  s = []                    # no capacity concept

## What main.go Shows

- Creating slices with literals, `make`, and slicing
- How `append` grows the backing array
- How sub-slices share the same backing array (mutation trap)

## Common Interview Traps

- A slice header is 3 fields: **pointer, length, capacity**
- Sub-slicing shares the backing array -- modifying one affects the other
- `append` returns a new slice header -- always reassign: `s = append(s, v)`
- When `len == cap`, `append` allocates a new, larger backing array
- `nil` slice and empty slice both have `len == 0` but `nil` slice is `== nil`
- You cannot compare slices with `==` (except to `nil`)

## What to Say in Interviews

- "A slice is a three-word struct: pointer to backing array, length, and capacity."
- "I always reassign the result of append because it may allocate a new backing array."
- "Sub-slices share memory, so I use copy when I need an independent slice."

## Run It

```bash
go run ./02_data_structures/02_slices
```

```bash
python ./02_data_structures/02_slices/main.py
```

## TL;DR (Interview Summary)

- Slice = pointer + length + capacity (3-word header)
- `make([]T, len, cap)` preallocates capacity to reduce allocations
- `append` may allocate a new array -- always reassign
- Sub-slices share the backing array -- mutations are visible
- Use `copy` to create an independent slice
- Nil slice (`var s []int`) vs empty slice (`[]int{}`) -- both len 0
- Cannot compare slices with `==`
