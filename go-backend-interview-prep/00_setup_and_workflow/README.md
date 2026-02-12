# 00 -- Setup and Workflow

Everything you need to write, build, and test Go code.
No deep dives here -- just enough to get moving.

---

## Installing Go

- Download the binary from the official Go site
- Extract to `/usr/local/go`
- Add `/usr/local/go/bin` to your `PATH`
- Verify: `go version`

On macOS with Homebrew:

```bash
brew install go
```

That's it. Go ships as a single binary with batteries included.

---

## go env Basics

`go env` prints all Go environment variables.

Key ones to know:

- `GOPATH` -- workspace root (default `~/go`), where downloaded modules live
- `GOROOT` -- where Go itself is installed
- `GOBIN` -- where `go install` puts binaries
- `GOOS` / `GOARCH` -- target OS and architecture (used for cross-compilation)
- `GOPROXY` -- module proxy URL (default `proxy.golang.org`)

```bash
# Print everything
go env

# Print one variable
go env GOPATH
```

---

## Modules: go mod

Go modules are how Go manages dependencies.
Every Go project starts with `go mod init`.

### go mod init

```bash
# Creates go.mod in the current directory
go mod init github.com/yourname/projectname
```

- `go.mod` tracks your module name and dependencies
- The module path is usually your repo URL

### go mod tidy

```bash
go mod tidy
```

- Adds missing dependencies
- Removes unused dependencies
- Run this after adding/removing imports

### go mod vendor (optional)

```bash
go mod vendor
```

- Copies all dependencies into a `vendor/` folder
- Useful for CI or air-gapped environments
- Most projects skip this and rely on the module cache

---

## Essential CLI Tools

### go fmt

```bash
go fmt ./...
```

- Auto-formats all Go files
- Not optional -- Go enforces one formatting style
- Run before every commit

### go vet

```bash
go vet ./...
```

- Catches suspicious code (e.g., unreachable code, bad printf args)
- Lightweight static analysis built into Go

### go test

```bash
go test ./...
```

- Runs all tests in the project
- Test files end in `_test.go`
- Test functions start with `Test` and take `*testing.T`

### golangci-lint (third-party, recommended)

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run ./...
```

- Runs many linters in one pass
- Highly configurable
- Common in production codebases

---

## Folder Convention for This Repo

```
01_go_basics/
  variables/
    main.go          <-- runnable example
    variables_test.go <-- tests (if applicable)
    README.md         <-- short explanation + interview notes
  constants/
    main.go
    README.md
```

- One subfolder per topic
- Each subfolder is a standalone `main` package (runnable)
- README in each subfolder: explanation, example walkthrough, pitfalls

---

## Package Naming

- Package name = last segment of the import path
- Keep names short, lowercase, single word when possible
- `package main` is special -- it is the entry point for executables
- Test files use the same package name (or `package foo_test` for black-box tests)

Good names: `http`, `json`, `auth`, `store`

Bad names: `httpHelpers`, `json_utils`, `MyPackage`

---

## TL;DR

- Install Go, verify with `go version`
- Start every project with `go mod init <module-path>`
- Use `go mod tidy` after changing imports
- Format with `go fmt`, lint with `go vet`, test with `go test`
- Each topic folder in this repo has a runnable `main.go` and a README
- Package names are short, lowercase, no underscores
