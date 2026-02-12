package main

import "fmt"

func main() {
	// --- Defer: LIFO order ---
	fmt.Println("--- Defer LIFO ---")
	fmt.Println("start")
	defer fmt.Println("deferred 1 (runs third)")
	defer fmt.Println("deferred 2 (runs second)")
	defer fmt.Println("deferred 3 (runs first)")
	fmt.Println("end")
	// Output order: start, end, deferred 3, deferred 2, deferred 1

	// --- Defer args evaluated immediately ---
	fmt.Println("\n--- Defer arg evaluation ---")
	x := 10
	defer fmt.Println("deferred x:", x) // captures x=10 NOW
	x = 99
	fmt.Println("current x:", x) // prints 99

	// --- Defer for cleanup pattern ---
	fmt.Println("\n--- Cleanup pattern ---")
	doWork()

	// --- Panic + Recover ---
	fmt.Println("\n--- Panic + Recover ---")
	safeCall()
	fmt.Println("program continues after recovered panic")

	// --- Common trap: recover must be in deferred func ---
	fmt.Println("\n--- Recover only works in defer ---")
	fmt.Println("This line proves the program did not crash.")
}

func doWork() {
	fmt.Println("  opening resource")
	defer fmt.Println("  closing resource (deferred)")

	fmt.Println("  doing work...")
	fmt.Println("  work done")
	// "closing resource" prints after this function returns
}

func safeCall() {
	// recover() must be inside a deferred function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  recovered from panic:", r)
		}
	}()

	fmt.Println("  about to panic")
	panic("something broke")
	// Code below never executes
}
