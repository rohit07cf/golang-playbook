# Interface Composition

## What It Is

- Go interfaces can **embed** other interfaces to compose larger ones
- **ELI10:** Composing interfaces is like stacking job requirements -- "must speak English AND drive" is two small interfaces combined.
- This follows the "small interfaces" philosophy

## Why It Matters

- Go's standard library uses this heavily (`io.ReadWriter` = `io.Reader` + `io.Writer`)
- Interviewers test whether you design with small, composable interfaces

## Syntax Cheat Sheet

```go
type Reader interface { Read(p []byte) (int, error) }
type Writer interface { Write(p []byte) (int, error) }

// Composed: embeds Reader + Writer
type ReadWriter interface {
    Reader
    Writer
}
```

**Go vs Python**

```
Go:  type ReadWriter interface { Reader; Writer }   // embedding
Py:  class ReadWriter(Reader, Writer, Protocol): ... # multiple inheritance
```

## What main.go + main.py Show

- Defining small interfaces and composing them
- A concrete type satisfying the composed interface
- Passing the concrete type where either sub-interface or composed interface is expected

## Common Interview Traps

- Composed interfaces require ALL embedded methods to be satisfied
- Prefer many small interfaces over one large interface
- A type satisfying `ReadWriter` also satisfies `Reader` and `Writer` individually
- You can embed any number of interfaces
- Name collisions between embedded interfaces cause a compile error

## What to Say in Interviews

- "I compose interfaces by embedding, keeping each interface small and focused."
- "io.ReadWriter is a classic example -- it embeds io.Reader and io.Writer."
- "A type satisfying a composed interface automatically satisfies each sub-interface."

## Run It

```bash
go run ./04_interfaces_and_generics/02_interface_composition
```

```bash
python ./04_interfaces_and_generics/02_interface_composition/main.py
```

## TL;DR (Interview Summary)

- Embed interfaces to compose larger ones
- `ReadWriter` = `Reader` + `Writer`
- Satisfying the composed interface satisfies each sub-interface
- Prefer many small interfaces (1-2 methods each)
- Name collisions cause compile errors
- Standard library uses this pattern extensively
