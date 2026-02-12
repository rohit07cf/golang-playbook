# 04 -- Interfaces and Generics

Polymorphism in Go: implicit interfaces and type parameters.
This module covers how Go achieves polymorphism without inheritance,
how interfaces compose, how type assertions work, how errors fit
the interface model, and how generics add type-safe reusable code.

Each example includes main.go + a Python equivalent main.py.

After this module you can:

- Define and implement interfaces (implicitly)
- Compose small interfaces into larger ones
- Use type assertions and type switches safely
- Understand errors as interfaces
- Write generic functions and types with constraints
- Decide when to use interfaces vs generics

---

## Subtopics

| Folder | What You Learn |
|--------|---------------|
| `01_interfaces_basics` | Implicit satisfaction, interface values, nil traps |
| `02_interface_composition` | Embedding interfaces, small vs large interfaces |
| `03_empty_interface_and_type_assertions` | `any`, comma-ok type assertion |
| `04_type_switch` | Dispatching on concrete type at runtime |
| `05_errors_as_interfaces` | The `error` interface, custom errors, wrapping |
| `06_stringer_and_custom_formatting` | `fmt.Stringer`, `fmt.Formatter` basics |
| `07_generics_basics` | Type parameters, basic generic functions and types |
| `08_generics_constraints` | `comparable`, `constraints` package, custom constraints |
| `09_generics_vs_interfaces` | When to use each, trade-offs, interview guidance |
| `_quick_revision` | Last-minute cheat sheet |

---

## 10-Min Revision Path

1. Skim `_quick_revision/README.md` for the full cheat sheet
2. Re-read `01_interfaces_basics` -- implicit satisfaction is the #1 question
3. Re-read `03_empty_interface_and_type_assertions` -- `any` and comma-ok form
4. Re-read `05_errors_as_interfaces` -- error handling is always asked
5. Re-read `07_generics_basics` -- know the syntax and when to use them
6. Re-read `09_generics_vs_interfaces` -- the "when do I use which?" question

---

## Common Interface/Generic Mistakes

- Declaring that a type "implements" an interface (Go has no `implements` keyword)
- Forgetting the comma-ok form of type assertion (single-value form panics on failure)
- Storing a nil pointer in an interface -- the interface itself is not nil
- Making interfaces too large (Go favors small, 1-2 method interfaces)
- Using `any` everywhere instead of specific interfaces or generics
- Overusing generics for code that only needs a simple interface
- Forgetting that generic type parameters are resolved at compile time
- Using `==` on type parameters without the `comparable` constraint

---

## TL;DR

- Interfaces are satisfied implicitly -- no `implements` keyword
- Prefer small interfaces (1-2 methods): `io.Reader`, `io.Writer`, `error`
- `any` is an alias for `interface{}` -- accepts everything, gives nothing
- Type assertions: always use comma-ok form to avoid panics
- `error` is an interface with one method: `Error() string`
- Generics add compile-time type safety for reusable code
- Use interfaces for behavior; use generics for type-safe containers/algorithms
- Nil pointer in interface != nil interface -- classic trap
