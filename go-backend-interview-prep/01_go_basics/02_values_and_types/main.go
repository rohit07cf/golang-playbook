package main

import "fmt"

func main() {
	// --- Integers ---
	var age int = 30
	var bigNum int64 = 9_000_000_000
	fmt.Println("int:", age)
	fmt.Println("int64:", bigNum)

	// --- Floats ---
	var pi float64 = 3.14159
	fmt.Println("float64:", pi)

	// --- Booleans ---
	var active bool = true
	fmt.Println("bool:", active)

	// --- Strings ---
	var greeting string = "hello, go"
	fmt.Println("string:", greeting)
	fmt.Println("string length (bytes):", len(greeting))

	// --- Byte vs Rune ---
	var b byte = 'A' // uint8
	var r rune = 'Z' // int32 (Unicode code point)
	fmt.Printf("byte: %c (%d)\n", b, b)
	fmt.Printf("rune: %c (%d)\n", r, r)

	// --- No implicit conversion ---
	// This would NOT compile:
	//   result := age + pi
	// You must cast explicitly:
	result := float64(age) + pi
	fmt.Println("explicit cast (int -> float64):", result)

	// --- Type inspection ---
	fmt.Printf("type of age: %T\n", age)
	fmt.Printf("type of pi: %T\n", pi)
	fmt.Printf("type of greeting: %T\n", greeting)
}
