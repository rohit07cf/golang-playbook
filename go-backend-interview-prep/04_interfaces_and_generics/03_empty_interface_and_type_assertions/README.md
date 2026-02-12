# Empty Interface and Type Assertions

## What It Is

- `any` (alias for `interface{}`) accepts a value of **any type**
- A **type assertion** extracts the concrete type from an interface value

## Why It Matters

- `any` is how Go handles truly generic containers (pre-generics era)
- Type assertions are the primary way to recover concrete types from interfaces

## Syntax Cheat Sheet

```go
// any accepts anything
func printVal(v any) { fmt.Println(v) }

// Type assertion (single-value form -- panics on failure!)
s := v.(string)

// Type assertion (comma-ok form -- safe)
s, ok := v.(string)
if ok { fmt.Println(s) }
```

**Go vs Python**

```
Go:  var v any = "hello"; s, ok := v.(string)
Py:  v: object = "hello"; s = v if isinstance(v, str) else None
```

## What main.go + main.py Show

- Storing different types in `any`
- Safe vs unsafe type assertion
- Why comma-ok form prevents panics

## Common Interview Traps

- Single-value assertion `v.(T)` **panics** if the type is wrong
- Always use comma-ok form: `val, ok := v.(T)`
- `any` gives up all type safety -- you must assert to get it back
- `any` is just `interface{}` with a new name (Go 1.18+)
- An uninitialized `any` is nil, not zero-valued

## What to Say in Interviews

- "I always use the comma-ok form of type assertion to avoid panics."
- "any is useful for truly polymorphic containers, but I prefer generics or specific interfaces when possible."
- "Since Go 1.18, any is an alias for interface{} -- they are identical."

## Run It

```bash
go run ./04_interfaces_and_generics/03_empty_interface_and_type_assertions
```

```bash
python ./04_interfaces_and_generics/03_empty_interface_and_type_assertions/main.py
```

## TL;DR (Interview Summary)

- `any` = `interface{}` -- accepts any type, provides no methods
- Type assertion: `val, ok := v.(T)` -- always use comma-ok form
- Single-value form `v.(T)` panics on wrong type
- `any` sacrifices type safety -- recover it with assertions
- Prefer specific interfaces or generics over `any`
- Uninitialized `any` is nil
