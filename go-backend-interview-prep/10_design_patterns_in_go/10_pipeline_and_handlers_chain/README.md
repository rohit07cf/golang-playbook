# Pipeline and Handlers Chain

## What It Is

- **Pipeline / Chain of Responsibility**: a message passes through a sequence of handlers; each can process, modify, or reject it
- **ELI10:** A pipeline is a factory assembly line -- each station does one job and passes the item to the next.
- In Go: each handler has the same interface and calls the next handler

## Why It Matters

- Pipelines separate concerns: validate -> enrich -> send, each step independent
- **ELI10:** Without a pipeline, one giant function does everything -- with a pipeline, each worker handles one step and hands it off.
- Interviewers test whether you can compose processing steps cleanly

## Syntax Cheat Sheet

```go
// Go: handler interface + chain
type Handler interface {
    Handle(msg *Message) error
}
type Chain struct { handlers []Handler }
func (c *Chain) Run(msg *Message) error {
    for _, h := range c.handlers {
        if err := h.Handle(msg); err != nil { return err }
    }
    return nil
}
```

```python
# Python: list of callables
class Pipeline:
    def __init__(self, *handlers):
        self.handlers = handlers
    def run(self, msg):
        for h in self.handlers:
            h.handle(msg)
```

> Both: message flows through handlers in order. Each handler does one thing.

## Tiny Example

- `main.go` -- validate -> enrich -> log -> send pipeline for notifications
- `main.py` -- same chain with duck-typed handler objects

## Common Interview Traps

- **Handler modifies message without copy**: downstream sees mutations -- be explicit
- **No error short-circuit**: if validation fails, skip remaining handlers
- **Tight coupling between steps**: each handler should work independently
- **Ordering bugs**: validate must come before send -- order matters

## What to Say in Interviews

- "I build a pipeline of handlers: validate, enrich, log, send -- each does one thing"
- "If any handler returns an error, the pipeline short-circuits -- no partial processing"
- "This is the same pattern as HTTP middleware but for business logic processing"

## Run It

```bash
go run ./10_design_patterns_in_go/10_pipeline_and_handlers_chain/
python ./10_design_patterns_in_go/10_pipeline_and_handlers_chain/main.py
```

## TL;DR (Interview Summary)

- Pipeline: message flows through ordered handlers
- Each handler does one thing: validate, enrich, log, send
- Short-circuit on error: don't process invalid messages
- Handler interface: `Handle(msg) error` -- composable and testable
- Same concept as HTTP middleware but for domain logic
- Order matters: validate before send, enrich before log
