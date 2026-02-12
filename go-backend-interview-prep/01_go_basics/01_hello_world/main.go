package main

import "fmt"

func main() {
	// The simplest Go program.
	// package main + func main() = executable entry point.
	fmt.Println("Hello, World!")

	// Print without a trailing newline
	fmt.Print("Go ")
	fmt.Print("is ")
	fmt.Println("compiled.")

	// Formatted output with Printf
	language := "Go"
	year := 2009
	fmt.Printf("%s was released in %d\n", language, year)
}
