package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// --- Example 1: os.Args ---
	fmt.Println("=== os.Args ===")
	fmt.Println("  program:", os.Args[0])
	fmt.Println("  all args:", os.Args[1:])

	// --- Example 2: flag package ---
	fmt.Println("\n=== flag package ===")
	name := flag.String("name", "world", "who to greet")
	count := flag.Int("count", 1, "how many times")
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	fmt.Printf("  name=%q count=%d verbose=%v\n", *name, *count, *verbose)
	fmt.Println("  remaining args:", flag.Args())

	for i := 0; i < *count; i++ {
		fmt.Printf("  hello, %s!\n", *name)
	}

	// --- Example 3: environment variables ---
	fmt.Println("\n=== Environment variables ===")

	// Getenv returns "" if unset
	home := os.Getenv("HOME")
	fmt.Println("  HOME:", home)

	// LookupEnv distinguishes unset from empty
	val, exists := os.LookupEnv("MY_APP_SECRET")
	if exists {
		fmt.Println("  MY_APP_SECRET:", val)
	} else {
		fmt.Println("  MY_APP_SECRET: (not set)")
	}

	// List env vars matching a prefix
	fmt.Println("\n=== Env vars starting with 'GO' ===")
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "GO") {
			fmt.Println(" ", env)
		}
	}
}
