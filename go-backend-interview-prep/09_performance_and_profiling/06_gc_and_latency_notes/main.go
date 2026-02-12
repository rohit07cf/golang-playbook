package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== GC and Latency Notes ===")
	fmt.Println()

	// Baseline
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	baseGC := m.NumGC
	fmt.Printf("  Baseline: NumGC=%d, HeapAlloc=%d KB\n\n", m.NumGC, m.HeapAlloc/1024)

	// Allocate lots of small objects to trigger GC
	fmt.Println("  Allocating 1M small objects...")
	start := time.Now()
	var sink []*[64]byte
	for i := 0; i < 1_000_000; i++ {
		obj := new([64]byte)
		sink = append(sink, obj)
	}
	allocTime := time.Since(start)

	runtime.ReadMemStats(&m)
	fmt.Printf("  After alloc: NumGC=%d (+%d), HeapAlloc=%d KB\n",
		m.NumGC, m.NumGC-baseGC, m.HeapAlloc/1024)
	fmt.Printf("  Alloc time: %s\n\n", allocTime)

	// Release references and force GC
	sink = nil
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("  After release + GC: HeapAlloc=%d KB\n", m.HeapAlloc/1024)
	fmt.Printf("  Total GC pause: %s\n", time.Duration(m.PauseTotalNs))
	fmt.Printf("  NumGC total: %d\n\n", m.NumGC)

	// Show recent pause times
	fmt.Println("  Recent GC pauses (most recent first):")
	for i := 0; i < 5; i++ {
		idx := (int(m.NumGC) - 1 - i) % 256
		if idx < 0 {
			break
		}
		pause := time.Duration(m.PauseNs[idx])
		fmt.Printf("    GC #%d: %s\n", int(m.NumGC)-i, pause)
	}

	fmt.Println()
	fmt.Println("--- Key GC concepts ---")
	fmt.Println("  Go GC: concurrent mark-and-sweep")
	fmt.Println("    - Tri-color algorithm (white/grey/black)")
	fmt.Println("    - Mostly concurrent -- very short STW pauses")
	fmt.Println("    - STW pauses typically < 1ms")
	fmt.Println()
	fmt.Println("  GOGC (GC trigger):")
	fmt.Println("    GOGC=100 (default): GC when heap = 2x live data")
	fmt.Println("    GOGC=200: GC when heap = 3x live data (fewer GCs, more memory)")
	fmt.Println("    GOGC=off: disable GC (careful!)")
	fmt.Println()
	fmt.Println("  Reduce GC pressure:")
	fmt.Println("    1. Preallocate slices: make([]T, 0, n)")
	fmt.Println("    2. Use value types (not pointers) in hot paths")
	fmt.Println("    3. sync.Pool for frequently allocated/freed objects")
	fmt.Println("    4. strings.Builder instead of += concat")
	fmt.Println("    5. Reduce allocation rate = fewer GC cycles")
}
