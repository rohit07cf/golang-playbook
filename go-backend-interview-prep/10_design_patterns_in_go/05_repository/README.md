# Repository Pattern

## What It Is

- **Repository**: an interface that abstracts data access (CRUD) behind a clean API
- Service layer calls the repository interface; implementation can be in-memory, SQL, or NoSQL

## Why It Matters

- Decouples business logic from storage -- swap DB without touching services
- Interviewers test whether you separate data access from business logic

## Syntax Cheat Sheet

```go
// Go: repository interface
type NotificationRepo interface {
    Save(n Notification) error
    FindByID(id string) (Notification, error)
    List() ([]Notification, error)
}
// In-memory implementation for tests
type MemoryRepo struct { data map[string]Notification }
```

```python
# Python: repository (duck typing or Protocol)
class MemoryRepo:
    def save(self, n: Notification) -> None: ...
    def find_by_id(self, id: str) -> Notification: ...
    def list_all(self) -> list[Notification]: ...
```

> Both: service depends on the interface; swap implementations freely.

## Tiny Example

- `main.go` -- NotificationRepo interface + MemoryRepo; service uses the interface
- `main.py` -- same with dict-backed repository

## Common Interview Traps

- **Putting SQL in handlers**: always go through a repository
- **Returning ORM models from repo**: return domain types, not DB types
- **No interface**: hardcoding the DB makes testing impossible
- **Too many methods**: keep repo focused on one entity's CRUD

## What to Say in Interviews

- "I define a repository interface in the service layer and inject the implementation"
- "For tests I use an in-memory repo; for production I use a SQL repo -- same interface"
- "The repo returns domain types, not ORM models -- keeping the boundary clean"

## Run It

```bash
go run ./10_design_patterns_in_go/05_repository/
python ./10_design_patterns_in_go/05_repository/main.py
```

## TL;DR (Interview Summary)

- Repository = interface for data access (Save, Find, List, Delete)
- Service depends on interface, not on DB driver
- In-memory repo for tests, SQL repo for production
- Return domain types from the repo, not DB-specific types
- Keep one repo per entity -- don't create a god repo
- Define the interface where it's consumed (service layer)
