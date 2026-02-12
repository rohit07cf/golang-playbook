package main

import (
	"encoding/json"
	"fmt"
)

// Struct tags control JSON field names and behavior.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"` // omit if empty string
	Age      int    `json:"age"`
	Password string `json:"-"` // never include in JSON
	Active   bool   `json:"active,omitempty"`
}

func main() {
	// === MARSHAL: struct -> JSON ===
	fmt.Println("--- Marshal ---")
	u := User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Age:      30,
		Password: "secret123",
		Active:   true,
	}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(data))
	// Password is excluded (-), all other fields present

	// === MARSHAL with omitempty ===
	fmt.Println("\n--- omitempty ---")
	u2 := User{Name: "Bob", Age: 25}
	// Email is "" -> omitted; Active is false -> omitted
	data2, _ := json.Marshal(u2)
	fmt.Println("JSON:", string(data2))

	// === PRETTY PRINT ===
	fmt.Println("\n--- MarshalIndent ---")
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))

	// === UNMARSHAL: JSON -> struct ===
	fmt.Println("--- Unmarshal ---")
	input := `{"name":"Charlie","email":"c@test.com","age":28,"active":true}`
	var u3 User
	err = json.Unmarshal([]byte(input), &u3) // must pass pointer
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Printf("parsed: %+v\n", u3)

	// === Unknown fields are silently ignored ===
	fmt.Println("\n--- Unknown fields ---")
	weird := `{"name":"Dave","unknown_field":"ignored","age":35}`
	var u4 User
	json.Unmarshal([]byte(weird), &u4)
	fmt.Printf("unknown fields ignored: %+v\n", u4)

	// === Partial JSON (missing fields get zero values) ===
	fmt.Println("\n--- Partial JSON ---")
	partial := `{"name":"Eve"}`
	var u5 User
	json.Unmarshal([]byte(partial), &u5)
	fmt.Printf("partial: %+v\n", u5) // Age=0, Email=""

	// === Marshal a slice ===
	fmt.Println("\n--- Slice of structs ---")
	team := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	teamJSON, _ := json.Marshal(team)
	fmt.Println("team:", string(teamJSON))

	// === Marshal a map ===
	fmt.Println("\n--- Map ---")
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}
	cfgJSON, _ := json.Marshal(config)
	fmt.Println("config:", string(cfgJSON))
}
