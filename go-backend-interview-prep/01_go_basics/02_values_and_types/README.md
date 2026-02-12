# Values and Types

## What It Is

- Go is statically typed -- every value has a fixed type at compile time
- Basic types: `int`, `float64`, `bool`, `string`, `byte`, `rune`

## Why It Matters

- Type mismatches are caught at compile time, not at runtime
- Understanding types prevents subtle bugs with numeric overflow and precision

## Syntax Cheat Sheet

```go
// Integers
var a int = 42          // platform-dependent size (usually 64-bit)
var b int64 = 42        // explicit 64-bit

// Floats
var f float64 = 3.14    // 64-bit float (default for literals)

// Booleans
var ok bool = true

// Strings (immutable, UTF-8 encoded)
var s string = "hello"

// Byte and Rune
var ch byte = 'A'       // alias for uint8
var r rune = 'Z'        // alias for int32 (Unicode code point)
```

## What main.go Shows

- Declaring and printing each basic type
- Demonstrating that Go does not implicitly convert between types

## Common Interview Traps

- `int` size is platform-dependent (32 or 64 bit) -- not fixed
- No implicit type conversion: `int + float64` does not compile
- Strings are immutable; you cannot change a character in place
- `byte` is just `uint8`; `rune` is just `int32`
- Default float literal is `float64`, not `float32`

## What to Say in Interviews

- "Go is statically typed with no implicit conversions -- you must cast explicitly."
- "Strings in Go are immutable byte slices, UTF-8 encoded by default."
- "I use int for general integers, and sized types like int64 only when the spec demands it."

## Run It

```bash
go run ./01_go_basics/02_values_and_types
```

## TL;DR (Interview Summary)

- Go basic types: `int`, `float64`, `bool`, `string`, `byte`, `rune`
- No implicit type conversion -- explicit casts required
- `int` size depends on platform (32 or 64 bit)
- Strings are immutable UTF-8 byte sequences
- `byte` = `uint8`, `rune` = `int32`
- Default numeric literal types: `int` for integers, `float64` for decimals
