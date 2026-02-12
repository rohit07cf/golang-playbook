# Hello World

## What It Is

- The smallest runnable Go program
- **ELI10:** Your first Go program is like a baby's first word -- it doesn't do much, but it proves everything is wired up correctly.
- Demonstrates the entry point: `package main` + `func main()`

## Why It Matters

- Every Go executable must have `package main` and `func main()`
- **ELI10:** Without `package main` and `func main()`, your program is like a car with no ignition -- all the parts are there, but nothing starts.
- Understanding this is step zero for writing any Go code

## Syntax Cheat Sheet

```go
package main          // declares an executable package

import "fmt"          // imports the formatting package

func main() {        // entry point -- no args, no return
    fmt.Println("hello")
}
```

**Go vs Python**
Go:  `fmt.Println("hello")`
Py:  `print("hello")`

## What main.go Shows

- A complete, minimal Go program that prints output
- How `import` brings in the standard library

## Common Interview Traps

- `main` must be in `package main` -- no exceptions
- `func main()` takes no arguments and returns nothing
- To access CLI args, use `os.Args` (not function params)
- Unused imports cause a compile error, not a warning
- `fmt.Println` adds a newline; `fmt.Print` does not

## What to Say in Interviews

- "Every Go binary starts at package main, func main."
- "Go enforces no unused imports at compile time, keeping code clean."
- "The standard library is imported by path, like fmt or os."

## Run It

```bash
go run ./01_go_basics/01_hello_world
```

```bash
python ./01_go_basics/01_hello_world/main.py
```

## TL;DR (Interview Summary)

- `package main` + `func main()` = executable entry point
- `import "fmt"` for formatted I/O
- `fmt.Println` prints with a newline
- Unused imports = compile error
- No args on `func main()`; use `os.Args` for CLI args
