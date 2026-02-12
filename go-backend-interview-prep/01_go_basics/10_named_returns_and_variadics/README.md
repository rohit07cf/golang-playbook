# Named Returns and Variadics

## What It Is

- **Named returns**: return values with names, initialized to zero values
- **ELI10:** Variadics are like an "all you can eat" buffet parameter -- pass one item, five items, or zero, and the function happily takes them all.
- **Variadics**: functions that accept a variable number of arguments (`...`)

## Why It Matters

- Named returns appear in standard library code -- you must read them
- **ELI10:** Named returns are like pre-labeled shipping boxes -- the labels are already on them, so at the end you just say "ship it" without specifying what goes where.
- Variadics are how `fmt.Println` and `append` work

## Syntax Cheat Sheet

```go
// Named return values
func split(total int) (half, remainder int) {
    half = total / 2
    remainder = total % 2
    return  // "naked return" -- returns half and remainder
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Passing a slice to a variadic
vals := []int{1, 2, 3}
sum(vals...)   // spread operator
```

**Go vs Python**
Go:  `func sum(nums ...int) int     // variadic`
Py:  `def sum_all(*nums: int) -> int  # *args`

## What main.go Shows

- Named returns with naked return
- Variadic functions called with individual args and with a spread slice

## Common Interview Traps

- Naked returns hurt readability -- use sparingly, only in short functions
- Named return variables are initialized to their zero values
- `...` goes after the type in the signature: `nums ...int`
- `...` goes after the slice when calling: `sum(vals...)`
- Variadic must be the **last** parameter
- An empty variadic call is valid: `sum()` returns the zero value loop result

## What to Say in Interviews

- "I use named returns mainly for documentation, and avoid naked returns in long functions."
- "Variadic params are syntactic sugar over slices -- inside the function, it is a slice."
- "You spread a slice into a variadic with the ... suffix on the argument."

## Run It

```bash
go run ./01_go_basics/10_named_returns_and_variadics
```

```bash
python ./01_go_basics/10_named_returns_and_variadics/main.py
```

## TL;DR (Interview Summary)

- Named returns: `func f() (x int, err error)` -- pre-declared, zero-valued
- Naked return: `return` with no args returns named values
- Avoid naked returns in long functions -- readability suffers
- Variadic: `func f(args ...int)` -- last param only
- Inside the function, variadic param is a `[]int` (slice)
- Spread a slice: `f(mySlice...)`
