package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Expression switch ---
	day := "Wednesday"

	switch day {
	case "Monday", "Tuesday":
		fmt.Println("early week")
	case "Wednesday":
		fmt.Println("midweek")
	case "Thursday", "Friday":
		fmt.Println("late week")
	default:
		fmt.Println("weekend")
	}

	// --- Tagless switch (no expression) ---
	// Works like a clean if/else chain
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("morning")
	case hour < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

	// --- Switch with init statement ---
	switch lang := "Go"; lang {
	case "Go":
		fmt.Println("compiled, statically typed")
	case "Python":
		fmt.Println("interpreted, dynamically typed")
	default:
		fmt.Println("unknown language")
	}

	// --- Explicit fallthrough ---
	// fallthrough is unconditional: it enters the next case
	// without checking its condition.
	fmt.Println("--- Fallthrough demo ---")
	n := 1
	switch n {
	case 1:
		fmt.Println("case 1 matched")
		fallthrough
	case 2:
		fmt.Println("case 2 entered (via fallthrough, not by matching)")
	case 3:
		fmt.Println("case 3 (not reached)")
	}

	// --- Multiple matches in one case ---
	char := 'e'
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf("%c is a vowel\n", char)
	default:
		fmt.Printf("%c is a consonant\n", char)
	}
}
