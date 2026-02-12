# Packages and Imports

## What It Is

- Go code is organized into **packages** -- every `.go` file declares one
- **Exported** names start with an uppercase letter; lowercase = unexported (private)

## Why It Matters

- Package design is central to Go project architecture
- The uppercase/lowercase visibility rule comes up in every interview

## Syntax Cheat Sheet

```go
// Declare a package
package mypackage

// Import a single package
import "fmt"

// Import multiple packages
import (
    "fmt"
    "strings"
    "math/rand"
)

// Alias an import
import r "math/rand"

// Blank import (side effects only)
import _ "net/http/pprof"

// Exported (public):   strings.ToUpper
// Unexported (private): strings.indexFunc
```

## What main.go Shows

- Importing standard library packages
- Using exported functions and constants
- Demonstrating the uppercase = exported rule

## Common Interview Traps

- Uppercase first letter = exported; lowercase = unexported (no `public`/`private` keywords)
- Unused imports are a compile error
- `import _` is for side effects (init functions run, nothing else)
- Package name is the last element of the import path: `"math/rand"` -> `rand.Intn()`
- Circular imports are not allowed -- the compiler rejects them

## What to Say in Interviews

- "Go uses capitalization for visibility -- uppercase is exported, lowercase is package-private."
- "There are no circular imports in Go; the compiler enforces a DAG."
- "I use blank imports only for side-effect packages like database drivers."

## Run It

```bash
go run ./01_go_basics/12_packages_and_imports
```

## TL;DR (Interview Summary)

- Every file declares `package name`
- Uppercase first letter = exported (public)
- Lowercase first letter = unexported (private to package)
- No `public`/`private` keywords
- Unused imports = compile error
- `import _` for side-effect-only packages
- Circular imports are forbidden
- Package name = last segment of import path
