# JSON Basics

## What It Is

- Go's `encoding/json` package marshals (struct -> JSON) and unmarshals (JSON -> struct)
- **Struct tags** control the JSON field names and behavior

## Why It Matters

- Every backend API reads and writes JSON
- Struct tags and `omitempty` are asked about in nearly every Go interview

## Syntax Cheat Sheet

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"age"`
    Pass  string `json:"-"`               // always omitted
}

// Marshal: struct -> JSON bytes
data, err := json.Marshal(user)

// Unmarshal: JSON bytes -> struct
var u User
err := json.Unmarshal(data, &u)

// Pretty print
data, _ := json.MarshalIndent(user, "", "  ")
```

## What main.go Shows

- Marshaling a struct to JSON with struct tags
- Unmarshaling JSON back into a struct
- `omitempty` behavior and the `-` tag
- Working with `json.MarshalIndent` for readable output

## Common Interview Traps

- Only **exported** fields (uppercase) are marshaled -- unexported fields are silently ignored
- `omitempty` omits zero values (0, "", false, nil) -- not "empty" in a custom sense
- The `-` tag means "never include this field" (useful for passwords, tokens)
- `json.Unmarshal` requires a **pointer** argument (`&u`, not `u`)
- Unknown JSON fields are silently ignored by default (no error)
- `json.Number` can avoid float64 precision loss for large integers

## What to Say in Interviews

- "I use struct tags to control JSON field names and omit empty values."
- "Only exported fields are marshaled, so I keep internal fields lowercase."
- "I use the - tag for sensitive fields like passwords to prevent accidental exposure."

## Run It

```bash
go run ./02_data_structures/08_json_basics
```

## TL;DR (Interview Summary)

- `json.Marshal` = struct to JSON; `json.Unmarshal` = JSON to struct
- Struct tags: `` `json:"name"` `` controls the JSON key
- `omitempty` skips zero values (0, "", false, nil)
- `-` tag excludes a field entirely
- Only exported (uppercase) fields are marshaled
- Unmarshal needs a pointer: `json.Unmarshal(data, &target)`
- Unknown JSON fields are silently ignored
