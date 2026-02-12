package main

import "fmt"

func main() {
	// --- C-style for loop ---
	fmt.Println("--- C-style ---")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- While-style (condition only) ---
	fmt.Println("--- While-style ---")
	n := 1
	for n < 50 {
		n *= 2
	}
	fmt.Println("n doubled to:", n)

	// --- Infinite loop with break ---
	fmt.Println("--- Infinite + break ---")
	counter := 0
	for {
		if counter >= 3 {
			break
		}
		fmt.Println("counter:", counter)
		counter++
	}

	// --- Continue (skip iteration) ---
	fmt.Println("--- Continue (skip even) ---")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- Range over a string ---
	// Returns byte index and rune (Unicode code point)
	fmt.Println("--- Range over string ---")
	for idx, ch := range "Go!" {
		fmt.Printf("index=%d char=%c\n", idx, ch)
	}

	// --- Ignoring index with _ ---
	fmt.Println("--- Ignore index ---")
	for _, ch := range "abc" {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// --- Nested loop with labeled break ---
	fmt.Println("--- Labeled break ---")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // breaks out of BOTH loops
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()
}
