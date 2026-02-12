# Paths and Directories

## What It Is

- `filepath.Join` builds OS-safe paths; `filepath.WalkDir` traverses directory trees
- `os.MkdirTemp` / `os.MkdirAll` create temporary and nested directories

## Why It Matters

- Hardcoded path separators break cross-platform code
- Interviewers test file traversal and temp directory cleanup patterns

## Syntax Cheat Sheet

```go
// Go: paths + walk
p := filepath.Join("data", "users", "file.txt")
filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
    fmt.Println(path); return nil
})
dir, _ := os.MkdirTemp("", "prefix")
defer os.RemoveAll(dir)
```

```python
# Python: pathlib + os.walk
from pathlib import Path
p = Path("data") / "users" / "file.txt"
for root, dirs, files in os.walk("."):
    print(root, files)
```

> **Python differs**: `pathlib.Path` is the modern approach; `os.path.join`
> is the older equivalent of `filepath.Join`.

## Tiny Example

- `main.go` -- filepath.Join, MkdirTemp, create files, WalkDir
- `main.py` -- pathlib, tempfile.mkdtemp, os.walk

## Common Interview Traps

- **Hardcoded `/` or `\`**: always use `filepath.Join` (Go) or `pathlib` (Python)
- **Not cleaning up temp dirs**: always `defer os.RemoveAll(dir)`
- **WalkDir vs Walk**: `WalkDir` (Go 1.16+) is more efficient -- use it
- **filepath vs path**: `filepath` is for OS paths; `path` is for URL/slash paths

## What to Say in Interviews

- "I use filepath.Join for all path construction to stay cross-platform"
- "I create temp directories with os.MkdirTemp and always defer RemoveAll"
- "WalkDir is preferred over Walk since Go 1.16 -- it avoids unnecessary os.Stat calls"

## Run It

```bash
go run ./07_io_files_networking/04_paths_and_dirs/
python ./07_io_files_networking/04_paths_and_dirs/main.py
```

## TL;DR (Interview Summary)

- `filepath.Join("a", "b", "c.txt")` -- OS-safe path construction
- `os.MkdirTemp("", "prefix")` + `defer os.RemoveAll(dir)` -- temp dir pattern
- `os.MkdirAll(path, perm)` -- create nested directories
- `filepath.WalkDir` -- traverse directory trees (Go 1.16+)
- `filepath.Ext`, `filepath.Base`, `filepath.Dir` -- path components
- Python: `pathlib.Path` (modern), `os.walk` (traversal), `tempfile.mkdtemp`
