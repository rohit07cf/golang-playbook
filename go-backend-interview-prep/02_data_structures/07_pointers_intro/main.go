package main

import "fmt"

type Config struct {
	Host    string
	Port    int
	Verbose bool
}

func main() {
	// --- Basics: & and * ---
	x := 42
	p := &x // p is *int, holds address of x

	fmt.Println("x:", x)
	fmt.Println("&x (address):", p)
	fmt.Println("*p (deref):", *p)

	// --- Modify through pointer ---
	*p = 100
	fmt.Println("x after *p = 100:", x) // 100

	// --- Pointer to a struct ---
	fmt.Println("\n--- Struct pointer ---")
	cfg := &Config{Host: "localhost", Port: 8080}
	fmt.Println("cfg.Host:", cfg.Host) // auto-deref (no ->)
	cfg.Port = 9090
	fmt.Println("cfg.Port:", cfg.Port)

	// --- Pass by pointer to mutate ---
	fmt.Println("\n--- Pass by pointer ---")
	val := 10
	fmt.Println("before:", val)
	increment(&val)
	fmt.Println("after increment:", val) // 11

	// --- Pass struct by pointer ---
	fmt.Println("\n--- Struct by pointer ---")
	user := Config{Host: "example.com", Port: 80}
	enableVerbose(&user)
	fmt.Printf("after enableVerbose: %+v\n", user)

	// --- Pass by value (does NOT mutate) ---
	fmt.Println("\n--- Pass by value (no mutation) ---")
	num := 10
	noChange(num)
	fmt.Println("num unchanged:", num)

	// --- new() allocates zero-value and returns pointer ---
	fmt.Println("\n--- new() ---")
	np := new(int)
	fmt.Println("*new(int):", *np) // 0
	*np = 55
	fmt.Println("after assignment:", *np)

	// --- Nil pointer ---
	fmt.Println("\n--- Nil pointer ---")
	var nilPtr *int
	fmt.Println("nilPtr:", nilPtr)             // <nil>
	fmt.Println("nilPtr == nil:", nilPtr == nil) // true
	// *nilPtr would PANIC: "invalid memory address"

	// --- Pointer comparison ---
	fmt.Println("\n--- Pointer comparison ---")
	a := 10
	b := 10
	pa := &a
	pb := &b
	fmt.Println("pa == pb:", pa == pb) // false (different addresses)
	fmt.Println("*pa == *pb:", *pa == *pb) // true (same value)

	// --- Local variable pointer is safe (escape analysis) ---
	fmt.Println("\n--- Escape analysis ---")
	cp := createConfig()
	fmt.Printf("returned pointer: %+v\n", *cp)
}

func increment(p *int) {
	*p++ // modifies the original
}

func enableVerbose(c *Config) {
	c.Verbose = true // modifies the original struct
}

func noChange(n int) {
	n = 999 // modifies local copy only
}

// Returning a pointer to a local variable is safe in Go.
// The compiler moves it to the heap via escape analysis.
func createConfig() *Config {
	c := Config{Host: "api.example.com", Port: 443}
	return &c
}
