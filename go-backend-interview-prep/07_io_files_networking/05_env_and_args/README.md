# Environment Variables and Command-Line Arguments

## What It Is

- `os.Args` gives raw command-line arguments; `flag` package parses typed flags
- **ELI10:** Environment variables are the sticky notes on your monitor -- the program reads them on startup, no questions asked.
- `os.Getenv` / `os.LookupEnv` read environment variables

## Why It Matters

- Every backend service reads config from env vars or CLI flags
- **ELI10:** CLI args are what you shout at the program when you launch it; env vars are what the room already knows before it walks in.
- Interviewers ask about 12-factor app config (env vars > files)

## Syntax Cheat Sheet

```go
// Go: args + flags + env
args := os.Args[1:]                 // skip program name
name := flag.String("name", "world", "who to greet")
flag.Parse()
val := os.Getenv("HOME")
```

```python
# Python: sys.argv + argparse + os.environ
import sys, os, argparse
args = sys.argv[1:]
parser = argparse.ArgumentParser()
parser.add_argument("--name", default="world")
val = os.environ.get("HOME", "")
```

> **Python differs**: `argparse` is more feature-rich than Go's `flag`.
> Go's `flag` only supports `-flag value` syntax (no GNU-style `--flag=value`).

## Tiny Example

- `main.go` -- os.Args, flag parsing, os.Getenv, os.LookupEnv
- `main.py` -- sys.argv, argparse, os.environ

## Common Interview Traps

- **os.Args[0] is the program name**: actual args start at index 1
- **flag.Parse() must be called**: forgetting it means all flags are defaults
- **Getenv returns ""**: can't distinguish "not set" from "set to empty" -- use LookupEnv
- **flag doesn't support --**: Go's flag package uses single dash only (`-name`, not `--name`)

## What to Say in Interviews

- "I use environment variables for secrets and runtime config, following 12-factor principles"
- "os.LookupEnv distinguishes between unset and empty; os.Getenv cannot"
- "For complex CLI tools I'd use a third-party package, but flag covers basics"

## Run It

```bash
go run ./07_io_files_networking/05_env_and_args/ -name=Gopher extra1 extra2
python ./07_io_files_networking/05_env_and_args/main.py --name Gopher extra1 extra2
```

## TL;DR (Interview Summary)

- `os.Args` = raw args; `os.Args[0]` = program name
- `flag.String/Int/Bool` + `flag.Parse()` = typed flag parsing
- `os.Getenv("KEY")` = read env var (returns "" if unset)
- `os.LookupEnv("KEY")` = read env var + check if it exists
- 12-factor: use env vars for config, not files
- Python: `sys.argv`, `argparse`, `os.environ`
