package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== HTTP Performance Basics ===")
	fmt.Println()

	// Start a local test server
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "pong"})
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	server := &http.Server{Handler: mux}
	go server.Serve(ln)
	defer server.Close()

	base := "http://" + ln.Addr().String()
	url := base + "/ping"
	n := 100

	// --- Approach 1: new client per request (no connection reuse) ---
	fmt.Printf("  Making %d requests...\n\n", n)

	start := time.Now()
	for i := 0; i < n; i++ {
		// BAD: new transport per request = no connection reuse
		client := &http.Client{
			Timeout:   5 * time.Second,
			Transport: &http.Transport{DisableKeepAlives: true},
		}
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		io.ReadAll(resp.Body)
		resp.Body.Close()
	}
	noReuseTime := time.Since(start)

	// --- Approach 2: reused client (connection pooling) ---
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	start = time.Now()
	for i := 0; i < n; i++ {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		io.ReadAll(resp.Body)
		resp.Body.Close() // MUST close to return connection to pool
	}
	reuseTime := time.Since(start)

	fmt.Printf("  No reuse (new client each time): %s\n", noReuseTime)
	fmt.Printf("  Reused client (keep-alive):      %s\n", reuseTime)
	fmt.Printf("  Speedup: %.1fx\n\n", float64(noReuseTime)/float64(reuseTime))

	// --- Show what happens if you don't close body ---
	fmt.Println("--- Common mistake: not closing resp.Body ---")
	fmt.Println("  If you don't close the body, the connection is NOT")
	fmt.Println("  returned to the pool. After MaxIdleConnsPerHost")
	fmt.Println("  connections leak, all new requests open fresh TCP.")
	fmt.Println()

	fmt.Println("--- Production http.Client template ---")
	fmt.Println(`  client := &http.Client{`)
	fmt.Println(`      Timeout: 10 * time.Second,`)
	fmt.Println(`      Transport: &http.Transport{`)
	fmt.Println(`          MaxIdleConns:        100,`)
	fmt.Println(`          MaxIdleConnsPerHost: 10,`)
	fmt.Println(`          IdleConnTimeout:     90 * time.Second,`)
	fmt.Println(`      },`)
	fmt.Println(`  }`)
	fmt.Println()
	fmt.Println("  Key: create once at startup, reuse everywhere.")
}
