# Escape Analysis and Pointers

## What It Is

- **Escape analysis**: the Go compiler decides whether a variable lives on the stack (fast) or heap (slower, GC'd)
- A variable "escapes to heap" when it outlives its function scope (e.g., returned as a pointer)

## Why It Matters

- Stack allocation is nearly free; heap allocation requires GC -- understanding this is key to perf tuning
- Interviewers ask "when does a value escape to heap?" to test your understanding of Go internals

## Syntax Cheat Sheet

```go
// Go: check escape analysis
// go build -gcflags="-m" ./path/

func noEscape() int {
    x := 42       // stays on stack
    return x
}

func escapes() *int {
    x := 42       // escapes to heap (pointer returned)
    return &x
}
```

```python
# Python: everything is on the heap
# No escape analysis -- all objects are heap-allocated
# Python uses reference counting + GC instead
x = 42  # heap-allocated int object (always)
```

> **Key difference**: Go has stack vs heap choice; Python always uses heap + refcount.

## Tiny Example

- `main.go` -- shows functions that escape vs don't escape, with timing comparison
- `main.py` -- explains Python's heap-only model and reference counting

## Common Interview Traps

- **Assuming pointers are always faster**: returning a pointer forces heap allocation
- **Not checking escape analysis**: `go build -gcflags="-m"` is the diagnostic tool
- **Interface conversions cause escapes**: passing a value to an `interface{}` parameter escapes it
- **Closures capture by reference**: captured variables may escape
- **Slice growth escapes**: when a slice outgrows its backing array, new array goes to heap

## What to Say in Interviews

- "I check escape analysis with `go build -gcflags='-m'` to see what moves to heap"
- "Returning a pointer forces heap allocation; returning a value keeps it on the stack"
- "I prefer value types in hot paths to avoid heap allocations and GC pressure"

## Run It

```bash
go run ./09_performance_and_profiling/05_escape_analysis_and_pointers/
python ./09_performance_and_profiling/05_escape_analysis_and_pointers/main.py

# Check escape analysis output:
go build -gcflags="-m" ./09_performance_and_profiling/05_escape_analysis_and_pointers/ 2>&1
```

## TL;DR (Interview Summary)

- `go build -gcflags="-m"` shows escape decisions
- Returning a pointer = heap allocation (escaped)
- Returning a value = stack allocation (fast, no GC)
- Interface conversions often cause escapes
- Closures capturing variables can cause escapes
- Python: everything is heap-allocated -- no escape analysis
- Prefer value types in hot paths to reduce GC pressure
