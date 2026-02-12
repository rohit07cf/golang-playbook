# Arrays

## What It Is

- Fixed-size, ordered collection of elements of the same type
- The size is part of the type: `[3]int` and `[5]int` are different types

## Why It Matters

- Slices (the type you actually use daily) are built on top of arrays
- Understanding arrays makes slice behavior make sense

## Syntax Cheat Sheet

```go
// Declare with explicit size
var nums [3]int                // [0, 0, 0]

// Declare and initialize
colors := [3]string{"red", "green", "blue"}

// Let the compiler count
primes := [...]int{2, 3, 5, 7, 11}

// Access and set
nums[0] = 42
fmt.Println(nums[0])

// Length
fmt.Println(len(primes))      // 5
```

**Go vs Python**
Go:  var a [3]int              // fixed size, value type
Py:  a = [0, 0, 0]            # list, dynamic, reference type

## What main.go Shows

- Declaring arrays with different syntaxes
- Demonstrating that arrays are value types (assigning copies the data)

## Common Interview Traps

- Arrays are **value types** -- assigning or passing copies the entire array
- `[3]int` and `[4]int` are completely different types (cannot assign one to the other)
- You almost never use arrays directly -- slices are the standard
- `len()` works on arrays but the size is known at compile time
- There is no `append` for arrays (that is a slice operation)

## What to Say in Interviews

- "Arrays in Go are value types with a fixed size baked into the type."
- "In practice I use slices, but understanding arrays explains how slices work internally."
- "Passing an array to a function copies it; I use a slice or pointer to avoid that."

## Run It

```bash
go run ./02_data_structures/01_arrays
```

```bash
python ./02_data_structures/01_arrays/main.py
```

## TL;DR (Interview Summary)

- Arrays: fixed-size, value types -- size is part of the type
- `[3]int` != `[4]int` -- different types entirely
- Assigning an array copies all elements
- `[...]int{1,2,3}` lets the compiler count the size
- Use slices instead for almost everything
- Arrays explain *why* slices behave the way they do
