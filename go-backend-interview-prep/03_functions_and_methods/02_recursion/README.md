# Recursion

## What It Is

- A function that calls itself with a smaller subproblem
- Requires a **base case** to stop and a **recursive case** to continue

## Why It Matters

- Tree traversal, graph search, and divide-and-conquer all use recursion
- Interviewers test whether you can identify base cases and avoid stack overflow

## Syntax Cheat Sheet

```go
func factorial(n int) int {
    if n <= 1 {
        return 1     // base case
    }
    return n * factorial(n-1)  // recursive case
}
```

**Go vs Python**

```
Go:  func fib(n int) int { if n < 2 { return n }; return fib(n-1)+fib(n-2) }
Py:  def fib(n): return n if n < 2 else fib(n-1) + fib(n-2)
```

## What main.go Shows

- Factorial and Fibonacci (classic examples)
- Iterative vs recursive comparison
- Stack depth warning

## Common Interview Traps

- Go has **no tail-call optimization** -- deep recursion causes stack overflow
- Default goroutine stack starts small (a few KB) but grows dynamically up to ~1GB
- For large inputs, prefer an iterative solution or use an explicit stack
- Fibonacci without memoization is O(2^n) -- always mention this
- Mutual recursion (A calls B, B calls A) is valid but hard to reason about

## What to Say in Interviews

- "I start by identifying the base case, then the recursive reduction."
- "Go does not optimize tail calls, so I prefer iteration for large inputs."
- "For Fibonacci I would use memoization or an iterative approach in production."

## Run It

```bash
go run ./03_functions_and_methods/02_recursion
```

```bash
python ./03_functions_and_methods/02_recursion/main.py
```

## TL;DR (Interview Summary)

- Every recursion needs a base case and a recursive case
- Go has no tail-call optimization -- deep recursion can overflow
- Prefer iteration for large/unbounded inputs
- Naive Fibonacci is O(2^n) -- use memoization or iteration
- Default stack grows dynamically but has limits
- Identify base case first, then reduce the problem
