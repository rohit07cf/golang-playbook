package main

import "fmt"

// Package-level variables must use 'var' (not :=)
var appName string = "go-basics"

func main() {
	// --- var with explicit type ---
	var city string = "Tokyo"
	fmt.Println("city:", city)

	// --- var with type inference ---
	var population = 14_000_000
	fmt.Println("population:", population)

	// --- Short declaration (functions only) ---
	country := "Japan"
	fmt.Println("country:", country)

	// --- Multiple declarations ---
	var x, y int = 10, 20
	fmt.Println("x:", x, "y:", y)

	a, b := "hello", true
	fmt.Println("a:", a, "b:", b)

	// --- Zero values ---
	var zeroInt int
	var zeroStr string
	var zeroBool bool
	var zeroFloat float64
	fmt.Println("--- Zero Values ---")
	fmt.Printf("int:     %d\n", zeroInt)
	fmt.Printf("string:  %q\n", zeroStr) // %q shows quotes around empty string
	fmt.Printf("bool:    %t\n", zeroBool)
	fmt.Printf("float64: %f\n", zeroFloat)

	// --- Shadowing trap ---
	score := 100
	fmt.Println("outer score:", score)
	{
		// This creates a NEW variable in the inner scope
		score := 999
		fmt.Println("inner score:", score)
	}
	fmt.Println("outer score (unchanged):", score)

	// --- Package-level variable ---
	fmt.Println("app:", appName)

	// Uncommenting the line below causes a compile error:
	// "unused declared and not used"
	// unusedVar := 42
}
