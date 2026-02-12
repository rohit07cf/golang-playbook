package main

import (
	"encoding/json"
	"fmt"
)

// User demonstrates struct tags.
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"-"` // never included in JSON
}

// Address has nested struct.
type Profile struct {
	User    User   `json:"user"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"`
}

func main() {
	// --- Example 1: Marshal struct to JSON ---
	fmt.Println("=== Marshal ===")
	u := User{Name: "alice", Age: 30, Email: "alice@example.com", IsAdmin: true}
	data, _ := json.Marshal(u)
	fmt.Println("  json:", string(data))

	// --- Example 2: Pretty print ---
	fmt.Println("\n=== Pretty print ===")
	pretty, _ := json.MarshalIndent(u, "  ", "  ")
	fmt.Println(string(pretty))

	// --- Example 3: Unmarshal JSON to struct ---
	fmt.Println("\n=== Unmarshal ===")
	jsonStr := `{"name":"bob","age":25,"email":"bob@example.com"}`
	var u2 User
	json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Printf("  parsed: %+v\n", u2)
	fmt.Println("  IsAdmin (skipped):", u2.IsAdmin) // always false

	// --- Example 4: omitempty ---
	fmt.Println("\n=== omitempty ===")
	empty := User{Name: "charlie", Age: 0, Email: "c@x.com"}
	data, _ = json.Marshal(empty)
	fmt.Println("  age=0 omitted:", string(data))

	// --- Example 5: nested struct ---
	fmt.Println("\n=== Nested struct ===")
	p := Profile{
		User:    User{Name: "dana", Age: 28, Email: "dana@x.com"},
		City:    "NYC",
		Country: "",
	}
	data, _ = json.Marshal(p)
	fmt.Println("  json:", string(data))

	// --- Example 6: dynamic JSON with map ---
	fmt.Println("\n=== Dynamic JSON (map) ===")
	dynamic := map[string]any{
		"status": "ok",
		"count":  42,
		"tags":   []string{"go", "json"},
	}
	data, _ = json.Marshal(dynamic)
	fmt.Println("  json:", string(data))

	// Unmarshal into map
	var result map[string]any
	json.Unmarshal(data, &result)
	fmt.Printf("  parsed map: %v\n", result)
}
