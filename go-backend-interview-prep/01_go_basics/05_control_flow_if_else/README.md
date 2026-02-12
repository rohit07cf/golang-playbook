# Control Flow: If / Else

## What It Is

- Go's conditional branching -- similar to other languages
- Unique feature: an **init statement** can precede the condition

## Why It Matters

- The init-statement pattern is idiomatic Go (especially with error checks)
- Interviewers test whether you know about scoping in if-init

## Syntax Cheat Sheet

```go
// Basic if/else
if x > 0 {
    // positive
} else if x == 0 {
    // zero
} else {
    // negative
}

// If with init statement (variable scoped to if/else block)
if val := compute(); val > threshold {
    // val exists here
}
// val does NOT exist here
```

**Go vs Python**
Go:  `if val := f(); val > 0 { }  // init statement`
Py:  `if (val := f()) > 0: ...    # walrus operator (3.8+)`

## What main.go Shows

- Basic if/else chains
- If with init statements
- How the init variable is scoped to the if/else block

## Common Interview Traps

- No parentheses around conditions: `if (x > 0)` works but is not idiomatic
- Braces `{}` are always required, even for single-line bodies
- Variables declared in the init statement are scoped to the entire if/else chain
- There is no ternary operator in Go (`x ? a : b` does not exist)
- The init statement is separated by `;`, not `,`

## What to Say in Interviews

- "Go's if-init pattern keeps error-check variables scoped tightly."
- "There is no ternary in Go -- you always write an explicit if/else."
- "Braces are mandatory in Go, which prevents dangling-else bugs."

## Run It

```bash
go run ./01_go_basics/05_control_flow_if_else
```

```bash
python ./01_go_basics/05_control_flow_if_else/main.py
```

## TL;DR (Interview Summary)

- No parentheses needed around conditions
- Braces `{}` always required
- If-init statement: `if val := f(); val > 0 { ... }`
- Init variable is scoped to the if/else block only
- No ternary operator in Go
- Prefer if-init for keeping scope tight
