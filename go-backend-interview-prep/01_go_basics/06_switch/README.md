# Switch

## What It Is

- Go's multi-way branch statement
- **ELI10:** Switch is a vending machine -- you press one button and get exactly one item. No accidentally dispensing everything below it.
- Does **not** fall through by default (opposite of C/Java)

## Why It Matters

- Cleaner than long if/else chains
- **ELI10:** Without automatic fall-through, you will never accidentally order a sandwich AND a salad when you only wanted the sandwich.
- No-fallthrough default is a common interview question

## Syntax Cheat Sheet

```go
// Expression switch
switch day {
case "Mon", "Tue":
    fmt.Println("early week")
case "Fri":
    fmt.Println("friday")
default:
    fmt.Println("other")
}

// Switch with no expression (acts like if/else)
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}

// Explicit fallthrough
switch n {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}
```

**Go vs Python**
Go:  `switch x { case 1, 2: ... }     // no fallthrough`
Py:  `match x: case 1 | 2: ...        // structural pattern (3.10+)`

## What main.go Shows

- Expression switch with multiple values per case
- Tagless switch (no expression, boolean cases)
- Explicit `fallthrough` keyword

## Common Interview Traps

- Cases do NOT fall through -- each case breaks automatically
- `fallthrough` is explicit and transfers to the **next case body unconditionally**
- `fallthrough` does not re-evaluate the next case condition
- Multiple values in one case: `case "a", "b":` (comma-separated)
- Switch with no expression is a clean replacement for if/else chains

## What to Say in Interviews

- "Go switch does not fall through by default, which prevents bugs."
- "I use tagless switch as a cleaner alternative to long if/else chains."
- "fallthrough in Go is unconditional -- it always enters the next case."

## Run It

```bash
go run ./01_go_basics/06_switch
```

```bash
python ./01_go_basics/06_switch/main.py
```

## TL;DR (Interview Summary)

- No fallthrough by default (opposite of C/Java)
- Multiple values per case: `case "a", "b":`
- Tagless switch (`switch { ... }`) replaces if/else chains
- `fallthrough` is explicit and unconditional
- No need for `break` at end of cases
- Switch is the idiomatic way to branch on multiple conditions
