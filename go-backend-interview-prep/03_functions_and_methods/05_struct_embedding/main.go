package main

import "fmt"

// --- Base types ---
type Base struct {
	ID int
}

func (b Base) Describe() string {
	return fmt.Sprintf("Base{ID: %d}", b.ID)
}

type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

// --- Embedding: User embeds Base ---
type User struct {
	Base        // embedded (promoted fields and methods)
	Name string
}

// --- Embedding multiple structs ---
type Admin struct {
	User          // embeds User (which embeds Base)
	Logger        // embeds Logger
	Level  string
}

// --- Field shadowing ---
type Employee struct {
	Base
	ID   string // shadows Base.ID (different type!)
	Name string
}

func main() {
	// --- Basic embedding ---
	fmt.Println("--- Basic embedding ---")
	u := User{
		Base: Base{ID: 1},
		Name: "Alice",
	}
	fmt.Println("u.ID:", u.ID)             // promoted field
	fmt.Println("u.Describe():", u.Describe()) // promoted method
	fmt.Println("u.Name:", u.Name)

	// --- Deep embedding ---
	fmt.Println("\n--- Deep embedding ---")
	a := Admin{
		User:   User{Base: Base{ID: 42}, Name: "Bob"},
		Logger: Logger{Prefix: "ADMIN"},
		Level:  "super",
	}
	fmt.Println("a.ID:", a.ID)             // promoted through User -> Base
	fmt.Println("a.Name:", a.Name)         // promoted from User
	fmt.Println("a.Describe():", a.Describe()) // promoted from Base
	a.Log("system started")               // promoted from Logger

	// --- Field shadowing ---
	fmt.Println("\n--- Field shadowing ---")
	e := Employee{
		Base: Base{ID: 100},
		ID:   "EMP-001", // shadows Base.ID
		Name: "Charlie",
	}
	fmt.Println("e.ID:", e.ID)           // "EMP-001" (outer field wins)
	fmt.Println("e.Base.ID:", e.Base.ID) // 100 (access shadowed via type name)

	// --- Embedding is NOT inheritance ---
	fmt.Println("\n--- Receiver is inner type ---")
	// u.Describe() calls Base.Describe(), not User.Describe()
	// The receiver inside Describe() is Base, not User.
	fmt.Println("Describe sees Base only:", u.Describe())
}
