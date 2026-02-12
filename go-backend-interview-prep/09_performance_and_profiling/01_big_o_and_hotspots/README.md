# Big-O and Hotspots

## What It Is

- **Big-O notation**: describes how an algorithm's time/space grows with input size
- **ELI10:** Big-O is how fast your code gets worse as the input grows -- O(n2) is like inviting everyone to shake hands with everyone else.
- **Hotspot**: the one function or loop where your program spends most of its time

## Why It Matters

- Interviewers ask "what's the time complexity?" for nearly every coding question
- **ELI10:** Without knowing Big-O, you're guessing why your program is slow -- like blaming the oven when the recipe is wrong.
- Profiling always starts with finding the hotspot -- the 5% of code using 95% of time

## Syntax Cheat Sheet

```go
// Go: timing a section
start := time.Now()
result := slowFunc(data)
fmt.Println("elapsed:", time.Since(start))
```

```python
# Python: timing a section
import time
start = time.perf_counter()
result = slow_func(data)
print(f"elapsed: {time.perf_counter() - start:.4f}s")
```

> Both languages: wrap the suspect code, print the time, compare approaches.

## Tiny Example

- `main.go` -- O(n^2) nested-loop duplicate finder vs O(n) map-based approach, timed
- `main.py` -- same comparison with `time.perf_counter()`

## Common Interview Traps

- **Saying O(1) when it's O(n)**: map lookup is O(1) average, but iterating a map is O(n)
- **Ignoring constants**: O(n) with a huge constant can be slower than O(n log n) for small n
- **Forgetting space complexity**: the map-based approach trades memory for speed
- **Not measuring**: always time both approaches -- theory doesn't account for cache effects
- **Confusing best/average/worst**: quicksort is O(n log n) average but O(n^2) worst

## What to Say in Interviews

- "I'd start by identifying the hotspot -- the innermost loop or most-called function"
- "This nested loop is O(n^2); I can reduce it to O(n) with a hash map at the cost of O(n) space"
- "I always measure with benchmarks rather than guessing -- cache effects can surprise you"

## Run It

```bash
go run ./09_performance_and_profiling/01_big_o_and_hotspots/
python ./09_performance_and_profiling/01_big_o_and_hotspots/main.py
```

## TL;DR (Interview Summary)

- Big-O describes growth rate, not absolute speed
- O(1) < O(log n) < O(n) < O(n log n) < O(n^2) < O(2^n)
- Find the hotspot first -- don't optimize code that barely runs
- Map lookup is O(1) average -- use maps to eliminate nested loops
- Always state both time AND space complexity
- Measure with real data -- theory is a guide, not the answer
