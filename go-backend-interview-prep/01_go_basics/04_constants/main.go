package main

import "fmt"

// Simple constants
const appVersion = "1.0.0"
const maxConnections int = 100

// Enum-style with iota
const (
	StatusPending  = iota // 0
	StatusActive          // 1
	StatusInactive        // 2
	StatusDeleted         // 3
)

// iota with bit shifts (useful for file permissions, sizes, etc.)
const (
	KB = 1 << (10 * (iota + 1)) // 1 << 10 = 1024
	MB                           // 1 << 20 = 1048576
	GB                           // 1 << 30 = 1073741824
)

// iota resets in each new const block
const (
	Red   = iota // 0 (reset)
	Green        // 1
	Blue         // 2
)

func main() {
	// --- Simple constants ---
	fmt.Println("version:", appVersion)
	fmt.Println("max connections:", maxConnections)

	// --- Enum constants ---
	fmt.Println("--- Status Enum ---")
	fmt.Println("Pending:", StatusPending)
	fmt.Println("Active:", StatusActive)
	fmt.Println("Inactive:", StatusInactive)
	fmt.Println("Deleted:", StatusDeleted)

	// --- Bit-shift constants ---
	fmt.Println("--- Sizes ---")
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)

	// --- iota reset proof ---
	fmt.Println("--- Colors (iota reset) ---")
	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)

	// --- Untyped constant flexibility ---
	// 'pi' has no fixed type, so it works with float32 and float64
	const pi = 3.14159
	var f32 float32 = pi
	var f64 float64 = pi
	fmt.Println("float32 pi:", f32)
	fmt.Println("float64 pi:", f64)
}
