package main

import "fmt"

// --- Define a struct type ---
type User struct {
	Name  string
	Email string
	Age   int
}

// --- Factory function (Go convention for "constructors") ---
func NewUser(name, email string, age int) User {
	return User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func main() {
	// --- Create with field names ---
	u1 := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	fmt.Println("u1:", u1)

	// --- Zero-value struct ---
	var u2 User
	fmt.Printf("u2 (zero): %+v\n", u2) // %+v shows field names

	// --- Partial initialization ---
	u3 := User{Name: "Bob"}
	fmt.Printf("u3 (partial): %+v\n", u3)

	// --- Factory function ---
	u4 := NewUser("Charlie", "charlie@example.com", 25)
	fmt.Println("u4:", u4)

	// --- Access and modify fields ---
	u1.Age = 31
	fmt.Println("u1 after birthday:", u1.Age)

	// --- Structs are VALUE types ---
	fmt.Println("\n--- Value semantics ---")
	original := User{Name: "Dave", Age: 40}
	copied := original
	copied.Name = "Modified"
	fmt.Println("original:", original.Name) // "Dave" -- unchanged
	fmt.Println("copied:", copied.Name)     // "Modified"

	// --- Passing to a function copies ---
	fmt.Println("\n--- Pass by value ---")
	tryToModify(original)
	fmt.Println("after function:", original.Name) // still "Dave"

	// --- Anonymous struct (one-off use) ---
	fmt.Println("\n--- Anonymous struct ---")
	point := struct {
		X, Y int
	}{X: 10, Y: 20}
	fmt.Println("point:", point)

	// --- Struct comparison ---
	fmt.Println("\n--- Comparison ---")
	a := User{Name: "Eve", Age: 28}
	b := User{Name: "Eve", Age: 28}
	fmt.Println("a == b:", a == b) // true (all fields are comparable)

	// --- Printing with format verbs ---
	fmt.Println("\n--- Format verbs ---")
	fmt.Printf("%%v:  %v\n", u1)
	fmt.Printf("%%+v: %+v\n", u1) // with field names
	fmt.Printf("%%#v: %#v\n", u1) // Go syntax representation
}

func tryToModify(u User) {
	u.Name = "CHANGED" // modifies the copy, not the original
}
