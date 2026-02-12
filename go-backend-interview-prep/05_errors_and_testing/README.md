# 05 -- Errors and Testing

Go treats errors as **values**, not exceptions. This module covers
the entire error-handling story plus Go's built-in testing toolkit.

Each example includes `main.go` + Python equivalent `main.py`.

---

## What This Module Covers

- The `if err != nil` pattern and why Go chose it
- Wrapping / unwrapping errors (`%w`, `errors.Is`, `errors.As`)
- Custom error types vs sentinel errors
- When panic is appropriate (almost never)
- Defer-based cleanup patterns
- `go test` basics, table-driven tests, subtests, benchmarks
- Simple mocking via interfaces + fake implementations

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `01_error_basics` | `error` interface, `if err != nil`, returning errors |
| 02 | `02_wrap_errors_and_unwrap` | `fmt.Errorf` with `%w`, `errors.Is`, `errors.As` |
| 03 | `03_custom_error_types` | Implementing the `error` interface on structs |
| 04 | `04_sentinel_errors_vs_typed_errors` | `var ErrX = errors.New(...)` vs typed error structs |
| 05 | `05_panic_recover_vs_errors` | When to panic, when to return errors, recover basics |
| 06 | `06_defer_for_cleanup_patterns` | File close, mutex unlock, recover-in-defer |
| 07 | `07_testing_basics` | `go test`, `t.Error`, `t.Fatal`, test file conventions |
| 08 | `08_table_driven_tests` | Table-driven pattern, `t.Run` subtests |
| 09 | `09_mocking_and_dependency_injection_intro` | Interface-based fakes, dependency injection |
| 10 | `10_benchmarks_intro` | `testing.B`, `b.N`, running and reading benchmarks |

---

## 10-Min Revision Path

1. Skim `01_error_basics` -- recall the `if err != nil` pattern
2. Skim `02_wrap_errors_and_unwrap` -- `%w`, `errors.Is`, `errors.As`
3. Skim `04_sentinel_errors_vs_typed_errors` -- know when to use each
4. Skim `06_defer_for_cleanup_patterns` -- cleanup idioms
5. Skim `08_table_driven_tests` -- the Go testing pattern interviewers love
6. Skim `_quick_revision/` -- one-screen cheat sheet

---

## Common Errors / Testing Mistakes

- Forgetting to check `err` (the #1 Go bug)
- Using `%v` instead of `%w` when wrapping (loses the error chain)
- Comparing errors with `==` instead of `errors.Is`
- Using `panic` for expected failures (use `error` instead)
- Not closing resources in defer (file handles, DB connections)
- Writing tests without subtests (hard to debug which case failed)
- Benchmarking with compiler-optimized-away results
- Ignoring `go vet` and `go test -race` in CI

---

## TL;DR

- Go errors are **values** that implement `interface { Error() string }`
- Always check `err != nil` before using the success value
- Wrap with `fmt.Errorf("context: %w", err)` to preserve the chain
- Use `errors.Is` for sentinel checks, `errors.As` for typed checks
- Reserve `panic` for unrecoverable programmer bugs, not user errors
- Table-driven tests + `t.Run` = idiomatic Go testing
- Benchmarks use `testing.B` and run with `go test -bench=.`
- Python equivalent: exceptions, `raise ... from`, `unittest`, `timeit`
