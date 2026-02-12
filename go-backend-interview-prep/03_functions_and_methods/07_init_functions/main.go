package main

import "fmt"

// Package-level variable: initialized first.
var config = loadConfig()

func loadConfig() string {
	fmt.Println("1. package-level var initialized")
	return "production"
}

// First init function.
func init() {
	fmt.Println("2. first init() called")
	fmt.Println("   config =", config)
}

// Second init function (yes, multiple init() in one file is valid).
func init() {
	fmt.Println("3. second init() called")
}

func main() {
	fmt.Println("4. main() called")
	fmt.Println("   config =", config)

	// Execution order:
	// 1. Package-level variables (loadConfig)
	// 2. First init()
	// 3. Second init()
	// 4. main()

	// You CANNOT call init() explicitly:
	// init()  // compile error: undefined: init

	fmt.Println("\nOrder: pkg vars -> init() (in file order) -> main()")
}
