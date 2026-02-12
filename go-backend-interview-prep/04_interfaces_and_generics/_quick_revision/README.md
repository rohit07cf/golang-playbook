# Interfaces & Generics -- Quick Revision

> One-screen cheat sheet. Skim before interviews.

---

## 1. Interface Satisfaction (Implicit)

```go
// Go: no "implements" keyword -- just define the methods
type Saver interface { Save() error }

type File struct{}
func (f File) Save() error { return nil } // File satisfies Saver automatically
```

```python
# Python: Protocol = structural typing (same idea)
from typing import Protocol

class Saver(Protocol):
    def save(self) -> None: ...

class File:
    def save(self) -> None: ...  # satisfies Saver automatically
```

## 2. Interface Composition

```go
type Reader interface { Read(p []byte) (int, error) }
type Writer interface { Write(p []byte) (int, error) }
type ReadWriter interface { Reader; Writer } // embed both
```

```python
class Reader(Protocol):
    def read(self, p: bytes) -> int: ...

class Writer(Protocol):
    def write(self, p: bytes) -> int: ...

class ReadWriter(Reader, Writer, Protocol): ...  # multiple inheritance
```

## 3. any + Type Assertions

```go
func describe(val any) {
    s, ok := val.(string) // comma-ok form (safe)
    if ok { fmt.Println("string:", s) }
}
```

```python
def describe(val: object) -> None:
    if isinstance(val, str):  # isinstance = type assertion
        print("string:", val)
```

## 4. Type Switch

```go
switch v := val.(type) {
case int:    fmt.Println("int:", v)
case string: fmt.Println("str:", v)
default:     fmt.Println("other")
}
```

```python
match val:
    case int():   print("int:", val)
    case str():   print("str:", val)
    case _:       print("other")
```

## 5. Errors as Interfaces (Go) vs Exceptions (Python)

```go
// Go: error is an interface { Error() string }
type NotFoundError struct{ ID int }
func (e *NotFoundError) Error() string { return fmt.Sprintf("not found: %d", e.ID) }

// Wrap: fmt.Errorf("context: %w", err)
// Check: errors.Is(err, target), errors.As(err, &target)
```

```python
# Python: exceptions are classes
class NotFoundError(Exception):
    def __init__(self, id: int):
        super().__init__(f"not found: {id}")

# Wrap: raise NewError("context") from original_err
# Check: except NotFoundError as e:
```

## 6. Generics Basics + Constraints

```go
func Map[T any, U any](s []T, fn func(T) U) []U { ... }     // any = no constraint
func Sum[T ~int | ~float64](s []T) T { ... }                  // custom constraint
func Max[T cmp.Ordered](a, b T) T { ... }                     // stdlib constraint
func Index[T comparable](s []T, target T) int { ... }         // == and !=
```

```python
from typing import TypeVar
T = TypeVar("T")
Num = TypeVar("Num", int, float)

def map_fn(s: list[T], fn) -> list: ...    # any type
def sum_all(s: list[Num]) -> Num: ...      # constrained to int/float
# Python generics are hints only -- not enforced at runtime
```

## 7. When to Use Generic vs Interface

```go
// INTERFACE: behavior differs between types
type Store interface { GetUser(id int) string }  // DB, Mock, Cache...

// GENERIC: same logic, types differ
func Reverse[T any](s []T) []T { ... }           // works on any slice
```

```python
# Protocol: behavior differs
class Store(Protocol):
    def get_user(self, id: int) -> str: ...

# Generic: same logic, types differ
def reverse(s: list[T]) -> list[T]: ...
```

---

## TL;DR

| Concept | Go | Python |
|---|---|---|
| Interface | implicit satisfaction | Protocol / ABC |
| Composition | embed interfaces | multiple inheritance |
| Type check | `val.(Type)` / type switch | `isinstance()` / `match` |
| Error | `error` interface, wrap with `%w` | exceptions, `raise ... from` |
| Generic func | `func F[T constraint](...)` | `def f(x: T) -> T:` (hint only) |
| Constraint | `comparable`, `cmp.Ordered`, custom | `TypeVar` bounds (not enforced) |
| Rule of thumb | interface = behavior, generic = algorithm | duck typing handles most cases |
