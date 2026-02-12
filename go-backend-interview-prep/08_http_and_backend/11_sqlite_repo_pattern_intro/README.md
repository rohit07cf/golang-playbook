# SQLite / Repository Pattern Intro

## What It Is

- **Repository pattern**: an interface that abstracts data access; handlers depend on the interface, not the storage engine
- **ELI10:** The repository pattern is a storage valet -- your service says "park this car" without caring if the garage uses SQLite, Postgres, or a pile of sticky notes.
- Go: define a `Repository` interface, implement with in-memory map (or real DB via external driver)
- Python: `sqlite3` is in the stdlib -- can demonstrate a real SQLite repo

## Why It Matters

- Separating data access from business logic makes code testable and swappable
- **ELI10:** If your handler talks directly to the database, swapping databases means rewriting everything. The repo pattern lets you swap the engine without rebuilding the car.
- Interviewers ask: "how would you swap from SQLite to Postgres?" -- answer: swap the repo implementation

## Syntax Cheat Sheet

```go
// Go: repository interface
type ItemRepo interface {
    Create(item Item) (Item, error)
    GetByID(id int) (Item, error)
    List() ([]Item, error)
    Delete(id int) error
}

// In-memory implementation
type MemoryRepo struct {
    mu    sync.RWMutex
    items map[int]Item
}
```

```python
# Python: sqlite3 repository
import sqlite3

class SQLiteRepo:
    def __init__(self, db_path=":memory:"):
        self.conn = sqlite3.connect(db_path)
        self.conn.execute("CREATE TABLE IF NOT EXISTS items (...)")

    def create(self, name, price):
        cur = self.conn.execute("INSERT INTO items ...", (name, price))
        return cur.lastrowid
```

> **Go note**: Go's stdlib has `database/sql` but no SQLite driver built in.
> Real SQLite requires `github.com/mattn/go-sqlite3` (CGO) or `modernc.org/sqlite`.
> We demonstrate the pattern with an in-memory map implementation.

## Tiny Example

- `main.go` -- Repository interface + MemoryRepo + service + HTTP handlers (layered architecture)
- `main.py` -- SQLiteRepo using stdlib sqlite3 + HTTP handlers

## Common Interview Traps

- **No interface**: hardcoding storage in handlers makes testing impossible
- **Business logic in handlers**: handlers should call a service, not contain logic
- **Leaking implementation details**: handlers shouldn't know if it's SQLite or Postgres
- **Forgetting to close DB**: always `defer db.Close()`
- **Not using transactions**: multiple operations should be atomic

## What to Say in Interviews

- "I define a Repository interface and inject the implementation into handlers"
- "This lets me swap from in-memory to SQLite to Postgres without changing handler code"
- "I layer: handler -> service -> repository -- each layer has a single responsibility"

## Run It

```bash
go run ./08_http_and_backend/11_sqlite_repo_pattern_intro/
python ./08_http_and_backend/11_sqlite_repo_pattern_intro/main.py
```

## TL;DR (Interview Summary)

- Repository pattern: interface abstracts data access
- Handlers depend on interface, not concrete storage
- Go: `database/sql` + driver for real DBs; we use in-memory map for demo
- Python: `sqlite3` is stdlib -- can use real SQLite directly
- Layer: handler -> service -> repository
- Swap implementations without changing business logic
- Always close DB connections (`defer db.Close()`)
- Use transactions for multi-step operations
