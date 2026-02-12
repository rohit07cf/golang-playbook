# Strings, Bytes, and Builders

## What It Is

- **strings.Builder**: efficient way to build strings in Go -- writes to an internal buffer, one alloc
- **ELI10:** Concatenating strings in a loop is like rewriting the entire letter every time you add a word -- use strings.Builder instead.
- **[]byte vs string**: strings are immutable in Go; converting between them copies data

## Why It Matters

- String concatenation in loops is the #1 performance trap interviewers test
- **ELI10:** Knowing about Builder is the difference between copying a whole book per page vs just adding a page to a binder.
- Understanding `[]byte` vs `string` conversion cost prevents hidden allocations

## Syntax Cheat Sheet

```go
// Go: strings.Builder
var b strings.Builder
b.Grow(n)           // preallocate
b.WriteString("hi") // append
s := b.String()     // one final copy

// []byte -> string (copies!)
bs := []byte("hello")
s = string(bs)  // copy
```

```python
# Python: list + join
parts = []
parts.append("hi")
s = "".join(parts)    # one concat at end

# bytes vs str
b = b"hello"          # bytes literal
s = b.decode("utf-8") # bytes -> str
```

> **Go**: `strings.Builder` avoids repeated allocation. `+=` creates a new string each time.
> **Python**: `"".join(list)` is the idiomatic fast path. `+=` is optimized in CPython but not guaranteed.

## Tiny Example

- `main.go` -- benchmarks concat, Builder, and Join approaches with timing
- `main.py` -- benchmarks +=, join, and io.StringIO approaches with timing

## Common Interview Traps

- **Using += in a loop**: O(n^2) total copying in Go (and often in Python)
- **Forgetting Builder.Grow()**: without it, Builder reallocates as it grows
- **Unnecessary string/[]byte conversions**: each conversion copies the data
- **Using fmt.Sprintf in hot paths**: slower than strconv or direct Builder writes
- **Assuming Python += is always slow**: CPython has a refcount optimization, but it's fragile

## What to Say in Interviews

- "I use strings.Builder with Grow(n) for building strings in loops -- one allocation total"
- "String-to-byte conversion copies the data; I avoid unnecessary conversions in hot paths"
- "In Python, ''.join(list) is the equivalent of Go's strings.Builder"

## Run It

```bash
go run ./09_performance_and_profiling/07_strings_bytes_and_builders/
python ./09_performance_and_profiling/07_strings_bytes_and_builders/main.py
```

## TL;DR (Interview Summary)

- `strings.Builder` + `Grow(n)` -- O(n) string building, one allocation
- `+=` concat -- O(n^2) total work (new string per iteration)
- `strings.Join(slice, sep)` -- good when you already have a slice
- `[]byte` <-> `string` conversion copies data -- avoid in hot paths
- `fmt.Sprintf` is convenient but slow; use `strconv` in hot paths
- Python: `"".join(list)` is the fast path; `+=` is fragile optimization
- Measure with `-benchmem` to see allocs/op difference
