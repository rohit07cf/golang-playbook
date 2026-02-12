# 05 -- Errors and Testing

How Go handles failure and how you verify correctness.

---

## Planned Topics

- [ ] Errors -- the `error` interface, `errors.New`, `fmt.Errorf`
- [ ] Custom error types -- implementing the `error` interface
- [ ] Error wrapping -- `%w`, `errors.Is`, `errors.As`
- [ ] Sentinel errors -- `io.EOF`, when to use them
- [ ] Panic and recover -- when (almost never) to use them
- [ ] Basic testing -- `*testing.T`, `go test`
- [ ] Table-driven tests -- the Go testing pattern
- [ ] Test helpers -- `t.Helper()`, `t.Cleanup()`
- [ ] Benchmarks -- `*testing.B`, `go test -bench`
- [ ] Subtests -- `t.Run()`, organizing test cases
