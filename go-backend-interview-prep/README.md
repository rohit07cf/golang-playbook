# Go Backend Interview Prep

A structured, low-cognitive-load repository for learning Go
from zero to interview-ready backend engineer.

Each folder covers one domain. Each topic inside has a short README,
a runnable example, and interview notes with common pitfalls.
Built for people who have limited time and need practical results.

---

## How to Use This Repo (~30 min/day)

- **Pick one folder** in order (start at `00_`, then `01_`, etc.)
- **Read the README** for that topic (5 min)
- **Run the example** and tweak it (10 min)
- **Review the interview notes** at the bottom of each topic (5 min)
- **Write the example from scratch** without looking (10 min)
- Move to the next topic tomorrow

If you only have 15 minutes: read the README + run the example.

---

## Learning Path

Work through folders in order. Each builds on the previous.

```
00  Setup and Workflow        -- tooling, project layout
01  Go Basics                 -- syntax, types, control flow
02  Data Structures           -- slices, maps, structs, pointers
03  Functions and Methods     -- signatures, receivers, closures
04  Interfaces and Generics   -- polymorphism, type constraints
05  Errors and Testing        -- error handling, table tests
06  Concurrency               -- goroutines, channels, patterns
07  IO, Files, Networking     -- readers, writers, TCP/UDP
08  HTTP and Backend          -- servers, middleware, JSON APIs
09  Performance and Profiling -- benchmarks, pprof, memory
10  Design Patterns in Go     -- options, DI, adapters
11  System Design in Go       -- architecture, trade-offs
12  Capstone Projects         -- full implementations
```

---

## Rules of This Repo

- Every concept has a runnable `.go` file
- READMEs are bullet-heavy, never paragraph-heavy
- No paragraph longer than 3 lines
- Interview pitfalls are marked clearly
- No fluff, no overengineering

---

## How to Run Examples

```bash
# Run a single file
go run ./01_go_basics/variables/main.go

# Run a folder with a main package
go run ./06_concurrency/goroutines/
```

## How to Test Examples

```bash
# Test everything
go test ./...

# Test one folder
go test ./05_errors_and_testing/table_tests/

# Verbose output
go test -v ./05_errors_and_testing/table_tests/
```

---

## 10 Minutes Before the Interview

Skim these in order:

1. `06_concurrency/README.md` -- goroutines, channels, select
2. `04_interfaces_and_generics/README.md` -- interface basics
3. `05_errors_and_testing/README.md` -- error wrapping, sentinel errors
4. `02_data_structures/README.md` -- slice internals, map gotchas
5. `08_http_and_backend/README.md` -- net/http patterns

These are the most commonly asked topics in Go backend interviews.

---

## Index

| Folder | What You Learn |
|--------|---------------|
| `00_setup_and_workflow` | Go installation, modules, tooling, project layout |
| `01_go_basics` | Variables, constants, types, loops, conditionals, strings |
| `02_data_structures` | Arrays, slices, maps, structs, pointers |
| `03_functions_and_methods` | Functions, multiple returns, closures, methods, receivers |
| `04_interfaces_and_generics` | Interfaces, type assertions, generics, type constraints |
| `05_errors_and_testing` | Error handling, custom errors, wrapping, table tests, benchmarks |
| `06_concurrency` | Goroutines, channels, select, context, sync primitives, patterns |
| `07_io_files_networking` | io.Reader/Writer, file ops, bufio, TCP/UDP basics |
| `08_http_and_backend` | net/http, routing, middleware, JSON encode/decode, REST patterns |
| `09_performance_and_profiling` | pprof, benchmarks, memory allocation, escape analysis |
| `10_design_patterns_in_go` | Functional options, dependency injection, adapters, interface segregation |
| `11_system_design_in_go` | Architecture decisions, trade-offs, scaling patterns |
| `12_capstone_projects` | HTTP API service, concurrent pipeline, CLI tool |
| `tools` | Helper scripts and shared utilities |

---

## Progress

- [x] `00_setup_and_workflow` -- Implemented
- [x] `01_go_basics` -- Implemented (12 topics + quick revision)
- [x] `02_data_structures` -- Implemented (8 topics + quick revision)
- [x] `03_functions_and_methods` -- Implemented (8 topics + quick revision)
- [x] `04_interfaces_and_generics` -- Implemented (9 topics + quick revision)
- [x] `05_errors_and_testing` -- Implemented (10 topics + quick revision)
- [x] `06_concurrency` -- Implemented (12 topics + quick revision)
- [x] `07_io_files_networking` -- Implemented (11 topics + quick revision)
- [x] `08_http_and_backend` -- Implemented (11 topics + quick revision)
- [x] `09_performance_and_profiling` -- Implemented (10 topics + quick revision)
- [ ] `10_design_patterns_in_go`
- [ ] `11_system_design_in_go`
- [ ] `12_capstone_projects`

---

## Planned Capstone Projects

- **HTTP API Service** -- CRUD with middleware, graceful shutdown, structured logging
- **Concurrent Pipeline** -- fan-out/fan-in data processing with context cancellation
- **CLI Tool** -- flag parsing, file I/O, exit codes, testable design
