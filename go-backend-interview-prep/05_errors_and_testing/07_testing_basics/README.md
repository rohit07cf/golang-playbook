# Testing Basics

## What It Is

- Go has a built-in testing framework: `testing` package + `go test` command
- **ELI10:** A test is a robot that runs your code and screams if something breaks -- so you don't have to check manually.
- Test files end in `_test.go`, test functions start with `Test` and take `*testing.T`

## Why It Matters

- Every Go project uses `go test` -- interviewers expect you to write tests fluently
- **ELI10:** Tests are your safety net -- refactor fearlessly because the robot catches your falls.
- No external frameworks needed (unlike most other languages)

## Syntax Cheat Sheet

```go
// Go: file must be *_test.go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    if got != 5 {
        t.Errorf("Add(2,3) = %d, want 5", got)
    }
}
```

```python
# Python: unittest
import unittest
class TestAdd(unittest.TestCase):
    def test_add(self):
        self.assertEqual(add(2, 3), 5)
```

## Tiny Example

- `main.go` -- defines `Add` and `Abs` functions
- `add_test.go` -- tests for `Add` and `Abs` using `t.Error`, `t.Fatal`, `t.Helper`
- `main.py` -- same functions in Python
- `add_test.py` -- unittest tests for the Python functions

## Common Interview Traps

- **t.Error vs t.Fatal**: `Error` continues the test; `Fatal` stops it immediately
- **Test file naming**: must end in `_test.go` or `go test` won't find it
- **Test function naming**: must start with `Test` and take `*testing.T`
- **Testing unexported functions**: test files in the same package can access unexported functions
- **Forgetting t.Helper()**: without it, error line numbers point to the helper, not the caller

## What to Say in Interviews

- "Go tests live in `_test.go` files in the same package, no external framework needed"
- "I use `t.Fatal` for setup failures that make continuing meaningless, `t.Error` for assertion checks"
- "`t.Helper()` marks a function as a test helper so error locations are reported correctly"

## Run It

```bash
go run ./05_errors_and_testing/07_testing_basics/
python ./05_errors_and_testing/07_testing_basics/main.py
go test ./05_errors_and_testing/07_testing_basics/
python -m pytest ./05_errors_and_testing/07_testing_basics/add_test.py
```

## TL;DR (Interview Summary)

- Test files: `*_test.go`, functions: `TestXxx(t *testing.T)`
- `t.Error` / `t.Errorf`: report failure, keep running
- `t.Fatal` / `t.Fatalf`: report failure, stop this test
- `t.Helper()`: mark helper functions for correct line reporting
- Run: `go test ./path/to/package/`
- Python: `unittest.TestCase` with `self.assertEqual`, `self.assertRaises`
