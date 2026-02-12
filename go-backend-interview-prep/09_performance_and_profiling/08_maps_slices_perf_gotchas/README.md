# Maps and Slices -- Performance Gotchas

## What It Is

- **Slice preallocation**: using `make([]T, 0, n)` to avoid repeated growing and copying
- **ELI10:** Preallocating is like reserving seats at a restaurant -- you don't want the waiter moving your table three times as more friends arrive.
- **Map hints**: `make(map[K]V, n)` preallocates buckets; map growth is expensive

## Why It Matters

- Slice/map growth triggers allocation + copy -- avoidable with known sizes
- **ELI10:** Growing without preallocating is like buying a bigger box and repacking everything each time you get a new toy.
- Interviewers test whether you know to preallocate and understand growth mechanics

## Syntax Cheat Sheet

```go
// Go: preallocate slice
s := make([]int, 0, 1000)  // len=0, cap=1000
for i := 0; i < 1000; i++ {
    s = append(s, i)        // no realloc needed
}

// Go: preallocate map
m := make(map[string]int, 1000)  // hint: ~1000 entries
```

```python
# Python: list preallocation
s = [None] * 1000   # or use list comprehension
# dict: no size hint in stdlib
d = {}               # grows dynamically (well optimized)
```

> **Go**: explicit prealloc matters a lot for slices/maps.
> **Python**: CPython optimizes list/dict growth well, but prealloc still helps for large sizes.

## Tiny Example

- `main.go` -- benchmarks preallocated vs non-preallocated slices and maps
- `main.py` -- benchmarks list append vs preallocated list, dict patterns

## Common Interview Traps

- **Appending without capacity**: each grow copies the entire backing array
- **Slice growth factor**: Go roughly doubles capacity (1.25x for large slices)
- **Map iteration is random**: don't depend on map order
- **String map keys**: each lookup hashes the full string -- long keys are slow
- **Nil slice vs empty slice**: `var s []int` (nil) works with append; `s := []int{}` (empty) does too

## What to Say in Interviews

- "When I know the final size, I use make([]T, 0, n) to avoid reallocations"
- "Map growth rehashes all entries -- preallocating with make(map[K]V, n) helps"
- "For hot-path map lookups with string keys, shorter keys hash faster"

## Run It

```bash
go run ./09_performance_and_profiling/08_maps_slices_perf_gotchas/
python ./09_performance_and_profiling/08_maps_slices_perf_gotchas/main.py
```

## TL;DR (Interview Summary)

- `make([]T, 0, n)` -- preallocate slice when size is known
- `make(map[K]V, n)` -- hint for map bucket allocation
- Slice append without cap: O(n) total copies from repeated doubling
- Slice with cap: zero copies if n is correct
- Map growth rehashes everything -- expensive for large maps
- Nil slice works with append -- no need for `make` if starting empty
- Python: list/dict growth is well-optimized, but prealloc still helps at scale
