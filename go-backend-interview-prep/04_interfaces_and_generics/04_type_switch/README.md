# Type Switch

## What It Is

- A `switch` statement that branches on the **concrete type** of an interface value
- **ELI10:** A type switch is like airport customs -- "What are you? Passport? String? Int? Step this way."
- Cleaner than chaining multiple type assertions

## Why It Matters

- The idiomatic way to handle `any` or interface values with multiple possible types
- Interviewers expect you to know this over repeated `if val, ok := v.(T)` chains

## Syntax Cheat Sheet

```go
switch v := val.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Printf("unknown: %T\n", v)
}
```

**Go vs Python**

```
Go:  switch v := val.(type) { case string: ... }
Py:  match v: case str(): ...   # or isinstance chains
```

## What main.go + main.py Show

- Type switch on an `any` value with multiple cases
- The `default` case for unmatched types
- Multiple types in a single case

## Common Interview Traps

- `.(type)` only works inside a `switch` statement -- not in regular expressions
- The variable `v` in `switch v := val.(type)` has the concrete type in each case
- You can combine multiple types: `case int, float64:`
- If no case matches and there is no default, nothing happens (no panic)
- Type switch cannot be used with generics type parameters

## What to Say in Interviews

- "I use type switch instead of chained type assertions for cleaner code."
- "The variable in each case branch has the matched concrete type, not any."
- "Type switch is the idiomatic way to dispatch on multiple concrete types."

## Run It

```bash
go run ./04_interfaces_and_generics/04_type_switch
```

```bash
python ./04_interfaces_and_generics/04_type_switch/main.py
```

## TL;DR (Interview Summary)

- `switch v := val.(type)` branches on concrete type
- `v` has the concrete type in each case branch
- Multiple types per case: `case int, float64:`
- No match + no default = no-op (no panic)
- `.(type)` only valid inside switch
- Preferred over chained type assertions
