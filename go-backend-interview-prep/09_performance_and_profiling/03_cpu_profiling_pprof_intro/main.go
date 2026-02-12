package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// cpuIntensiveWork does repeated hashing to create a measurable CPU load.
func cpuIntensiveWork(iterations int) [32]byte {
	data := []byte("hello world performance profiling")
	var hash [32]byte
	for i := 0; i < iterations; i++ {
		hash = sha256.Sum256(data)
		data = hash[:]
	}
	return hash
}

// stringWork does string manipulation to show up as a separate hotspot.
func stringWork(iterations int) string {
	s := ""
	for i := 0; i < iterations; i++ {
		s += "x"
		if len(s) > 1000 {
			s = s[:100]
		}
	}
	return s
}

func main() {
	// --- Create CPU profile ---
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create profile:", err)
		return
	}
	defer f.Close()

	fmt.Println("=== CPU Profiling with pprof ===")
	fmt.Println("Starting CPU profile -> cpu.prof")

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// --- Run workloads ---
	start := time.Now()

	fmt.Println("\nRunning CPU-intensive hashing...")
	hash := cpuIntensiveWork(500_000)
	fmt.Printf("  hash result: %x... (%s)\n", hash[:4], time.Since(start))

	mid := time.Now()
	fmt.Println("Running string manipulation...")
	s := stringWork(100_000)
	fmt.Printf("  string len: %d (%s)\n", len(s), time.Since(mid))

	fmt.Printf("\nTotal: %s\n", time.Since(start))

	// Profile is written when StopCPUProfile runs (deferred)
	fmt.Println("\n--- How to analyze ---")
	fmt.Println("  go tool pprof cpu.prof")
	fmt.Println()
	fmt.Println("Key pprof commands:")
	fmt.Println("  top10         -- show top 10 functions by flat time")
	fmt.Println("  top -cum      -- sort by cumulative time")
	fmt.Println("  list funcName -- show annotated source for a function")
	fmt.Println("  web           -- open interactive SVG graph in browser")
	fmt.Println("  peek funcName -- show callers and callees")
	fmt.Println("  quit          -- exit pprof")
	fmt.Println()
	fmt.Println("For live services, import net/http/pprof:")
	fmt.Println("  import _ \"net/http/pprof\"")
	fmt.Println("  // Then: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10")
}
