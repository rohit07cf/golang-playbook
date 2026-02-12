# JSON Encode / Decode

## What It Is

- `json.Marshal` converts a Go struct to JSON bytes; `json.Unmarshal` does the reverse
- **ELI10:** JSON encoding is packing a suitcase with labels. Decoding is unpacking it and putting things in the right drawers.
- Struct tags (`json:"name"`) control field names, omit-empty, and skip behavior

## Why It Matters

- JSON is the lingua franca of web APIs -- every backend engineer marshals JSON daily
- **ELI10:** If APIs are people talking, JSON is the common language they all agreed to speak.
- Interviewers ask about struct tags, unexported fields, and omitempty behavior

## Syntax Cheat Sheet

```go
// Go: struct tags control JSON keys
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age,omitempty"`
    Admin bool   `json:"-"` // skip
}
data, _ := json.Marshal(user)
json.Unmarshal(data, &user)
```

```python
# Python: json + dataclass
from dataclasses import dataclass, asdict
import json
@dataclass
class User:
    name: str; age: int = 0
json.dumps(asdict(user))
json.loads(json_str)
```

> **Python differs**: no struct tags. Use `dataclass` + `asdict()` for
> structured marshal. Key naming must be handled manually if needed.

## Tiny Example

- `main.go` -- marshal/unmarshal structs, struct tags, pretty print, map[string]any
- `main.py` -- json.dumps/loads, dataclass, custom key mapping

## Common Interview Traps

- **Unexported fields are invisible to json**: lowercase fields won't marshal
- **omitempty skips zero values**: 0, "", false, nil are all omitted
- **`json:"-"` vs `json:"-,"`**: `-` skips the field; `"-,"` uses literal key `-`
- **Unmarshal into wrong type**: will silently skip fields that don't match
- **Encoder/Decoder for streams**: use `json.NewEncoder(w)` for HTTP responses, not `json.Marshal`

## What to Say in Interviews

- "I use struct tags to control JSON field names and omitempty for optional fields"
- "For HTTP handlers I use json.NewEncoder(w).Encode to stream directly to the response"
- "Unexported Go fields are invisible to encoding/json -- they must be capitalized"

## Run It

```bash
go run ./07_io_files_networking/06_json_encode_decode/
python ./07_io_files_networking/06_json_encode_decode/main.py
```

## TL;DR (Interview Summary)

- `json.Marshal(v)` -> `[]byte`; `json.Unmarshal(data, &v)` -> fills struct
- Struct tags: `json:"name,omitempty"` controls key name + empty behavior
- `json:"-"` skips the field entirely
- Unexported fields (lowercase) are **invisible** to json
- Use `json.NewEncoder(w)` for streaming (HTTP responses)
- `map[string]any` for dynamic JSON without a struct
- Python: `json.dumps()` / `json.loads()`, `dataclass` for structure
