# 01 -- Go Basics

Everything you need to write simple Go programs.
This module covers the syntax, types, control flow, and functions
that form the foundation of every Go program.

After this module you can:

- Declare variables and constants
- Use all control flow constructs (if, switch, for)
- Write functions with multiple returns and variadic args
- Understand defer, panic, and recover at a high level
- Organize code into packages

---

## Subtopics

| Folder | What You Learn |
|--------|---------------|
| `01_hello_world` | Entry point, `package main`, `fmt.Println` |
| `02_values_and_types` | Basic types: int, float64, bool, string |
| `03_variables` | `var`, `:=`, zero values, type inference |
| `04_constants` | `const`, `iota`, typed vs untyped constants |
| `05_control_flow_if_else` | Conditionals, init statements in if |
| `06_switch` | Expression switch, no fallthrough by default |
| `07_for_loops` | The only loop keyword in Go, three forms |
| `08_functions_intro` | Function signatures, parameters, returns |
| `09_multiple_return_values` | Returning (value, error), tuple-style returns |
| `10_named_returns_and_variadics` | Named return values, `...` variadic params |
| `11_defer_panic_recover_intro` | Deferred cleanup, crash handling basics |
| `12_packages_and_imports` | Import paths, package naming, visibility |
| `_quick_revision` | Last-minute interview cheat sheet |

---

## 10-Min Revision Path

1. Skim `_quick_revision/README.md` for the full cheat sheet
2. Re-read `03_variables` -- interviewers love zero-value questions
3. Re-read `07_for_loops` -- the only loop, know all three forms
4. Re-read `09_multiple_return_values` -- idiomatic error handling starts here
5. Re-read `11_defer_panic_recover_intro` -- common interview topic

---

## Common Go Beginner Mistakes

- Using `=` instead of `:=` for short declarations
- Forgetting that unused variables cause compile errors
- Expecting `switch` to fall through (it does not by default)
- Confusing `byte` (uint8) with `rune` (int32)
- Using `:=` to redeclare when you meant to reassign
- Forgetting that `for` is the only loop -- there is no `while`
- Ignoring the second return value (especially errors)

---

## TL;DR

- Go has a small, strict syntax -- the compiler enforces style
- Every Go program starts at `package main`, `func main()`
- Variables have zero values; nothing is uninitialized
- `:=` declares and assigns; `var` is for explicit typing or package-level
- `for` is the only loop; `switch` does not fall through
- Functions can return multiple values -- always check errors
- `defer` runs at function exit, LIFO order
- Unused variables and imports are compile errors, not warnings
