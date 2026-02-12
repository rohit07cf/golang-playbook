# For Loops

## What It Is

- `for` is the **only** loop keyword in Go
- Three forms: C-style, while-style, and infinite

## Why It Matters

- Interviewers expect you to know there is no `while` or `do-while`
- Understanding the three forms covers all looping needs

## Syntax Cheat Sheet

```go
// C-style: init; condition; post
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style: condition only
for n < 100 {
    n *= 2
}

// Infinite loop
for {
    // use break to exit
    break
}

// Range over a string
for index, char := range "hello" {
    fmt.Printf("%d: %c\n", index, char)
}
```

## What main.go Shows

- All three loop forms
- `break` and `continue`
- Ranging over strings (byte index + rune value)

## Common Interview Traps

- There is no `while` keyword -- use `for condition { }`
- `range` over a string gives byte index and rune (not byte)
- Forgetting `break` in an infinite loop = hung program
- Loop variable `i` is scoped to the for block
- `continue` skips to the next iteration, not out of the loop

## What to Say in Interviews

- "Go has only one loop keyword: for. It covers C-style, while, and infinite loops."
- "When I range over a string, I get byte indices and runes, not bytes."
- "I use the while-style for when I do not need an index or post statement."

## Run It

```bash
go run ./01_go_basics/07_for_loops
```

## TL;DR (Interview Summary)

- `for` is the only loop keyword (no while, no do-while)
- Three forms: C-style, while-style, infinite
- `break` exits the loop; `continue` skips to next iteration
- `range` over string: byte index + rune value
- Loop variables are scoped to the for block
- Infinite loop: `for { ... }` with explicit `break`
