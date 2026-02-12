package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	// --- Example 1: DNS lookup ---
	fmt.Println("=== DNS Lookup ===")
	hosts := []string{"localhost", "example.com"}

	for _, host := range hosts {
		addrs, err := net.LookupHost(host)
		if err != nil {
			fmt.Printf("  %s: error: %v\n", host, err)
		} else {
			fmt.Printf("  %s -> %v\n", host, addrs)
		}
	}

	// --- Example 2: HTTP client timeout ---
	fmt.Println("\n=== HTTP Client Timeout ===")

	// Local server: one fast, one slow endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fast response"))
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("slow response"))
	})

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	defer ln.Close()
	go http.Serve(ln, mux)
	base := "http://" + ln.Addr().String()

	// Fast request succeeds
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(base + "/fast")
	if err != nil {
		fmt.Println("  fast: error:", err)
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("  fast:", string(body))
	}

	// Slow request times out
	_, err = client.Get(base + "/slow")
	if err != nil {
		fmt.Println("  slow: timed out (expected):", err)
	}

	// --- Example 3: context timeout per request ---
	fmt.Println("\n=== Context Timeout (per-request) ===")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", base+"/slow", nil)
	noTimeoutClient := &http.Client{} // no global timeout
	_, err = noTimeoutClient.Do(req)
	if err != nil {
		fmt.Println("  context timeout (expected):", err)
	}

	// --- Example 4: dial timeout ---
	fmt.Println("\n=== Dial Timeout ===")
	// Try connecting to a non-routable address to trigger timeout
	start := time.Now()
	_, err = net.DialTimeout("tcp", "192.0.2.1:12345", 500*time.Millisecond)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("  dial timeout after %v: %v\n", elapsed.Round(time.Millisecond), err)
	}
}
