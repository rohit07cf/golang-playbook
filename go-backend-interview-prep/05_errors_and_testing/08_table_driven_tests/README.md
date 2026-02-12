# Table-Driven Tests

## What It Is

- A testing pattern where test cases are defined as a **slice of structs** and iterated with `t.Run`
- Each struct holds inputs, expected output, and a descriptive name

## Why It Matters

- This is the **idiomatic Go testing pattern** -- interviewers expect you to use it
- Makes adding cases trivial (just add a row) and failures easy to identify

## Syntax Cheat Sheet

```go
// Go: table-driven with t.Run subtests
tests := []struct{
    name string; input string; want int; wantErr bool
}{
    {"valid", "42", 42, false},
    {"negative", "-1", -1, false},
    {"bad", "abc", 0, true},
}
for _, tc := range tests {
    t.Run(tc.name, func(t *testing.T) { ... })
}
```

```python
# Python: subTest for the same pattern
cases = [("valid", "42", 42, False), ("bad", "abc", 0, True)]
for name, inp, want, want_err in cases:
    with self.subTest(name=name):
        ...
```

## Tiny Example

- `main.go` -- defines `ParseIntSafe` that wraps `strconv.Atoi` with a default
- `parse_test.go` -- table-driven tests with `t.Run` subtests
- `main.py` -- same function in Python
- `parse_test.py` -- unittest with `subTest`

## Common Interview Traps

- **Forgetting `t.Run`**: without subtests, a failure doesn't tell you which case
- **Sharing loop variable in goroutine**: `tc` is reused each iteration -- capture it (`tc := tc`)
- **Not naming test cases**: unnamed cases make failures cryptic
- **Skipping error cases**: always test both success and failure paths
- **Testing too much at once**: each case should test one behavior

## What to Say in Interviews

- "I use table-driven tests with `t.Run` so each case is a named subtest"
- "Adding a new case is just adding a row to the table -- no new function needed"
- "Failed subtests show the case name, making debugging straightforward"

## Run It

```bash
go run ./05_errors_and_testing/08_table_driven_tests/
python ./05_errors_and_testing/08_table_driven_tests/main.py
go test -v ./05_errors_and_testing/08_table_driven_tests/
python -m pytest ./05_errors_and_testing/08_table_driven_tests/parse_test.py
```

## TL;DR (Interview Summary)

- Define cases as `[]struct{ name, input, want, wantErr }`
- Iterate with `for _, tc := range tests { t.Run(tc.name, ...) }`
- Each subtest runs independently, reports its own name on failure
- Adding cases = adding rows, not writing new functions
- Python: `self.subTest(name=...)` gives the same effect
- This is THE pattern interviewers expect for Go tests
