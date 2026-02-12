# Errors & Testing -- Quick Revision

> One-screen cheat sheet. Skim before interviews.

---

## 1. Errors vs Panic

```go
// Go: errors for expected failures
if err != nil { return fmt.Errorf("context: %w", err) }

// panic for programmer bugs only
panic("unreachable code")
```

```python
# Python: exceptions for expected failures
raise ValueError("bad input")

# assert for programmer bugs
assert x > 0, "unreachable"
```

## 2. Wrap / Unwrap (Is / As)

```go
wrapped := fmt.Errorf("load: %w", ErrNotFound)   // wrap with %w
errors.Is(wrapped, ErrNotFound)                    // true (walks chain)
errors.As(wrapped, &typedErr)                      // extracts typed error
```

```python
raise RuntimeError("load") from original_err      # chain with 'from'
isinstance(e.__cause__, NotFoundError)             # walk chain manually
```

## 3. Sentinel vs Typed Errors

```go
var ErrNotFound = errors.New("not found")          // sentinel: simple, no data
type HTTPError struct{ Code int }                  // typed: carries data
func (e *HTTPError) Error() string { ... }
```

```python
class NotFoundError(Exception): pass               # sentinel-like
class HTTPError(Exception):                        # typed: has fields
    def __init__(self, code: int): self.code = code
```

## 4. Defer Cleanup

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()                                    // runs at function return, LIFO
```

```python
with open(path) as f:                              # context manager = defer
    data = f.read()
```

## 5. go test Patterns

```go
// Table-driven test with subtests
func TestParse(t *testing.T) {
    tests := []struct{ name, in string; want int }{
        {"ok", "42", 42}, {"bad", "x", 0},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) { ... })
    }
}
```

```python
# Python: subTest
for name, inp, want in cases:
    with self.subTest(name=name):
        self.assertEqual(parse(inp), want)
```

## 6. Benchmarks

```go
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ { Fib(20) }
}
// Run: go test -bench=. -benchmem
```

```python
import timeit
t = timeit.timeit(lambda: fib(20), number=1000)
print(f"{t/1000*1e6:.1f} us/op")
```

---

## Interview One-Liners

1. "Go errors are values -- `error` is `interface{ Error() string }`"
2. "Wrap with `%w`, check with `errors.Is` and `errors.As`"
3. "Panic is for programmer bugs; errors are for expected failures"
4. "Table-driven tests with `t.Run` = idiomatic Go testing"
5. "Benchmarks use `b.N` and run with `go test -bench=.`"
6. "I inject dependencies as interfaces so I can swap in fakes for testing"

---

## TL;DR

- `if err != nil` -- always check before using the result
- `%w` wraps, `%v` loses the chain
- `errors.Is` = sentinel check, `errors.As` = typed check
- Panic = bugs, error = expected failures
- `defer` for cleanup (LIFO), args evaluated at defer line
- Table-driven tests + `t.Run` = the pattern interviewers want
- Benchmarks: `testing.B`, `b.N`, `-benchmem`
- Python: exceptions + `raise...from` + `unittest` + `timeit`
