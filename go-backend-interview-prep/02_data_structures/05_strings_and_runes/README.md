# Strings and Runes

## What It Is

- Strings in Go are **immutable sequences of bytes**, UTF-8 encoded by default
- A `rune` is an alias for `int32` and represents a single Unicode code point

## Why It Matters

- Interviewers test whether you know `len()` returns bytes, not characters
- Correctly iterating over multi-byte characters requires understanding runes

## Syntax Cheat Sheet

```go
s := "Hello, world"

// Length in bytes
len(s)

// Index gives a byte, not a rune
s[0]                 // byte value

// Range iterates by rune (correct for Unicode)
for i, r := range s {
    // i = byte index, r = rune (int32)
}

// Convert between string, []byte, []rune
bytes := []byte(s)
runes := []rune(s)
back  := string(runes)

// Single rune from literal
r := 'A'             // type: rune (int32)
```

**Go vs Python**
Go:  len("Go言語")    // 8 (bytes)
Py:  len("Go言語")    # 4 (characters)

## What main.go Shows

- The difference between byte length and rune count
- Iterating by byte index vs iterating by rune
- Converting between strings, byte slices, and rune slices

## Common Interview Traps

- `len("cafe\u0301")` returns bytes (6), not visible characters (5 runes, 4 glyphs)
- Indexing a string `s[i]` gives a **byte**, not a rune
- `range` over a string yields **(byte index, rune)** -- index may jump
- Strings are immutable -- to modify, convert to `[]byte` or `[]rune`, change, convert back
- String concatenation in a loop is O(n^2) -- use `strings.Builder`
- `"a"` is a string literal; `'a'` is a rune literal

## What to Say in Interviews

- "Go strings are UTF-8 byte slices, so len returns bytes, not characters."
- "I use range to iterate runes correctly, since multi-byte characters span multiple indices."
- "For building strings in a loop I use strings.Builder to avoid O(n^2) allocation."

## Run It

```bash
go run ./02_data_structures/05_strings_and_runes
```

```bash
python ./02_data_structures/05_strings_and_runes/main.py
```

## TL;DR (Interview Summary)

- Strings are immutable UTF-8 byte slices
- `len(s)` = byte count, not character count
- `s[i]` = byte at position i, not a rune
- `range s` iterates by rune: `(byte_index, rune)`
- `[]rune(s)` gives the character slice; `len([]rune(s))` = character count
- Use `strings.Builder` for efficient string concatenation
- `'x'` is a rune; `"x"` is a string
