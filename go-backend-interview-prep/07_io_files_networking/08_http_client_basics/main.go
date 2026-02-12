package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	// --- Spin up a tiny local HTTP server ---
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom", "demo-value")
		fmt.Fprintf(w, "Hello from local server! Method=%s", r.Method)
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","count":42}`))
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // intentionally slow
		w.Write([]byte("slow response"))
	})

	// Listen on a random port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	baseURL := fmt.Sprintf("http://%s", listener.Addr().String())
	go http.Serve(listener, mux)
	defer listener.Close()

	fmt.Println("local server at", baseURL)

	// --- Example 1: GET with timeout ---
	fmt.Println("\n=== GET /hello ===")
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(baseURL + "/hello")
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("  status:", resp.StatusCode)
	fmt.Println("  body:", string(body))
	fmt.Println("  X-Custom:", resp.Header.Get("X-Custom"))

	// --- Example 2: GET JSON ---
	fmt.Println("\n=== GET /json ===")
	resp2, _ := client.Get(baseURL + "/json")
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println("  Content-Type:", resp2.Header.Get("Content-Type"))
	fmt.Println("  body:", string(body2))

	// --- Example 3: custom request with headers ---
	fmt.Println("\n=== Custom request with headers ===")
	req, _ := http.NewRequest("GET", baseURL+"/hello", nil)
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("User-Agent", "go-interview-prep/1.0")

	resp3, _ := client.Do(req)
	defer resp3.Body.Close()
	body3, _ := io.ReadAll(resp3.Body)
	fmt.Println("  body:", string(body3))

	// --- Example 4: timeout on slow endpoint ---
	fmt.Println("\n=== Timeout (1s on /slow) ===")
	fastClient := &http.Client{Timeout: 1 * time.Second}
	_, err = fastClient.Get(baseURL + "/slow")
	if err != nil {
		fmt.Println("  expected timeout:", err)
	}
}
