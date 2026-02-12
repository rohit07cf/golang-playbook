# 03 -- Functions and Methods

How Go defines behavior and attaches it to types.
This module covers closures, recursion, methods with receivers,
method sets, struct embedding, defer patterns, and init functions.

Each example includes main.go + a Python equivalent main.py.

After this module you can:

- Write closures and understand variable capture
- Define methods on types with value and pointer receivers
- Know which receiver type satisfies which interface (method sets)
- Use struct embedding for composition
- Use defer correctly in production code
- Understand init function execution order

---

## Subtopics

| Folder | What You Learn |
|--------|---------------|
| `01_closures` | Anonymous functions, captured variables, closure traps |
| `02_recursion` | Base cases, stack depth, tail recursion (Go has none) |
| `03_methods_and_receivers` | Value receivers vs pointer receivers |
| `04_method_sets` | Which receiver satisfies which interface |
| `05_struct_embedding` | Composition, promoted methods, field shadowing |
| `06_defer_deep_dive` | LIFO, loop traps, argument evaluation |
| `07_init_functions` | `func init()`, execution order, side effects |
| `08_first_class_functions` | Function types, callbacks, strategy pattern |
| `_quick_revision` | Last-minute cheat sheet |

---

## 10-Min Revision Path

1. Skim `_quick_revision/README.md` for the cheat sheet
2. Re-read `03_methods_and_receivers` -- value vs pointer receiver is asked constantly
3. Re-read `04_method_sets` -- key to understanding interface satisfaction
4. Re-read `01_closures` -- the loop-variable trap is a classic gotcha
5. Re-read `06_defer_deep_dive` -- defer-in-loop pitfall comes up often

---

## Common Function/Method Mistakes

- Capturing loop variable in closure (all goroutines see same value)
- Using value receiver when you need to mutate the struct
- Forgetting that method sets differ for `T` vs `*T`
- Deferring inside a loop (resources not released until function exits)
- Assuming Go has tail-call optimization (it does not)
- Shadowing embedded fields by declaring same-name fields
- Relying on init function order across packages

---

## TL;DR

- Closures capture variables by reference -- beware loop vars
- Value receiver = copy; pointer receiver = mutation
- `*T` method set includes both value and pointer receiver methods
- `T` method set includes only value receiver methods
- Struct embedding promotes methods -- composition over inheritance
- `defer` runs at function exit, LIFO; args evaluated immediately
- `func init()` runs automatically at package load -- no explicit call
