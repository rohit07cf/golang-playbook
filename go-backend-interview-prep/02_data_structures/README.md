# 02 -- Data Structures

How Go organizes and stores data in memory.
This module covers the core data structures every Go program uses:
arrays, slices, maps, strings, structs, and pointers.

After this module you can:

- Use arrays and slices correctly (and know the difference)
- Manipulate maps with the comma-ok idiom
- Work with UTF-8 strings and runes
- Define structs and use pointers to avoid copying
- Marshal and unmarshal JSON with struct tags

---

## Subtopics

| Folder | What You Learn |
|--------|---------------|
| `01_arrays` | Fixed-size, value semantics, when to use them |
| `02_slices` | Length vs capacity, append, backing arrays |
| `03_slice_tricks` | Delete, insert, filter, copy, gotchas |
| `04_maps` | Hash maps, nil trap, comma-ok, iteration order |
| `05_strings_and_runes` | UTF-8 encoding, byte vs rune, iteration |
| `06_structs_intro` | Defining types, field access, zero values |
| `07_pointers_intro` | &, *, nil, pass-by-pointer vs pass-by-value |
| `08_json_basics` | Marshal, Unmarshal, struct tags, omitempty |
| `_quick_revision` | Last-minute cheat sheet for all DS topics |

---

## 10-Min Revision Path

1. Skim `_quick_revision/README.md` for the one-screen summary
2. Re-read `02_slices` -- slice internals are the #1 data structure interview topic
3. Re-read `03_slice_tricks` -- append gotchas trip up even experienced devs
4. Re-read `04_maps` -- nil map write panic is a classic trap
5. Re-read `05_strings_and_runes` -- UTF-8 iteration is always asked
6. Re-read `07_pointers_intro` -- pointer vs value semantics come up constantly

---

## Common DS Mistakes in Go

- Treating slices like arrays (they share backing storage)
- Writing to a nil map (panic at runtime)
- Assuming map iteration order is stable (it is random)
- Using `len()` on a string and expecting character count (it returns bytes)
- Forgetting that `append` may return a new backing array
- Comparing slices with `==` (does not compile; use `slices.Equal` or loop)
- Modifying a struct value in a range loop (it is a copy)
- Ignoring the second return from map lookup (the comma-ok idiom)

---

## TL;DR

- Arrays are fixed-size, value types -- rarely used directly
- Slices are the workhorse: dynamic, reference a backing array
- `append` may allocate a new backing array -- always reassign: `s = append(s, v)`
- Maps are unordered; nil map reads return zero, nil map writes panic
- Strings are immutable byte slices; use `range` to iterate runes correctly
- Structs are value types; use pointers to avoid large copies
- JSON uses struct tags: `` `json:"name,omitempty"` ``
- Know slice length vs capacity -- this is interview gold
