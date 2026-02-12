# Mutexes

## What It Is

- `sync.Mutex` provides **mutual exclusion** -- only one goroutine can hold the lock
- **ELI10:** A mutex is the bathroom key at a gas station -- only one person at a time, everyone else waits in line
- `sync.RWMutex` allows **multiple readers** OR **one writer** (read-write lock)

## Why It Matters

- Shared state without a mutex = data race (undefined behavior)
- Interviewers test whether you know mutex vs channel for protecting shared state

## Syntax Cheat Sheet

```go
// Go: Mutex
var mu sync.Mutex
mu.Lock()
counter++       // protected
mu.Unlock()

// RWMutex: multiple readers, one writer
var rw sync.RWMutex
rw.RLock()      // shared read lock
val := cache[key]
rw.RUnlock()
```

```python
# Python: threading.Lock / threading.RLock
import threading
lock = threading.Lock()
with lock:
    counter += 1  # protected
```

> **Python differs**: the GIL protects some operations on built-in types,
> but you still need `threading.Lock` for compound operations (read-modify-write).

## Tiny Example

- `main.go` -- unsafe counter (data race), safe counter with Mutex, RWMutex for read-heavy cache
- `main.py` -- same patterns with `threading.Lock`

## Common Interview Traps

- **Forgetting to unlock**: always use `defer mu.Unlock()` right after Lock
- **ELI10:** Forgetting to unlock a mutex is like walking away with the bathroom key -- everyone waits forever
- **Copying a mutex**: `sync.Mutex` must not be copied (pass by pointer)
- **Locking twice**: calling `Lock()` twice on the same Mutex in the same goroutine = deadlock
- **RWMutex abuse**: don't use RWMutex if writes are frequent (RLock overhead)
- **Mutex vs channel**: mutex = protect shared state; channel = communicate between goroutines

## What to Say in Interviews

- "I use defer Unlock immediately after Lock to prevent forgetting to unlock"
- "For read-heavy workloads I use RWMutex -- multiple goroutines can RLock concurrently"
- "I prefer channels for communication and mutexes for protecting shared state"

## Run It

```bash
go run ./06_concurrency/08_mutexes/
python ./06_concurrency/08_mutexes/main.py
```

## TL;DR (Interview Summary)

- `mu.Lock()` / `defer mu.Unlock()` -- always pair, always defer
- `sync.RWMutex`: `RLock` for reads (shared), `Lock` for writes (exclusive)
- Never copy a Mutex -- pass by pointer
- `go test -race` detects races from missing locks
- Mutex = protect state; channel = communicate
- Python: `threading.Lock()`, use `with lock:` for auto-release
