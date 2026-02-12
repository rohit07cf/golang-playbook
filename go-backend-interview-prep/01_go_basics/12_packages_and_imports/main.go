package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// --- Using the fmt package ---
	fmt.Println("--- fmt package ---")
	fmt.Println("Println adds a newline")
	fmt.Printf("Printf with formatting: %d\n", 42)

	// --- Using the strings package ---
	fmt.Println("\n--- strings package ---")
	msg := "Hello, Go World"

	// Exported functions start with uppercase
	fmt.Println("ToUpper:", strings.ToUpper(msg))
	fmt.Println("ToLower:", strings.ToLower(msg))
	fmt.Println("Contains 'Go':", strings.Contains(msg, "Go"))
	fmt.Println("Replace:", strings.ReplaceAll(msg, "Go", "Gopher"))
	fmt.Println("Split:", strings.Split(msg, " "))

	// --- Using the math package ---
	fmt.Println("\n--- math package ---")
	fmt.Println("Pi:", math.Pi)     // Exported constant (uppercase P)
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("Max(3,7):", math.Max(3, 7))

	// --- Visibility rule ---
	fmt.Println("\n--- Visibility rule ---")
	fmt.Println("math.Pi is exported (uppercase P)")
	// math.pi would NOT compile -- lowercase = unexported

	// --- Package name = last segment ---
	// "math" package -> math.Pi
	// "strings" package -> strings.ToUpper
	// "fmt" package -> fmt.Println
	fmt.Println("\nPackage names match the last import path segment.")

	// --- Internal functions demonstrate visibility ---
	fmt.Println("\n--- Local visibility ---")
	fmt.Println("exported call:", ExportedGreet("Alice"))
	fmt.Println("unexported call:", unexportedHelper())
}

// ExportedGreet is visible outside this package (uppercase E)
func ExportedGreet(name string) string {
	return "Hello, " + name
}

// unexportedHelper is only visible within this package (lowercase u)
func unexportedHelper() string {
	return "I am package-private"
}
