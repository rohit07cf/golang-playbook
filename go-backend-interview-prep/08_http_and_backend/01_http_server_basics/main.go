package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello from Go HTTP server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)

	// Use a random port for demo; in production use a fixed port
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		// Port 8080 busy -- pick random
		listener, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
	}

	addr := listener.Addr().String()
	fmt.Printf("Server running on http://%s\n", addr)
	fmt.Println("  GET /hello  -> greeting")
	fmt.Println("  GET /health -> JSON health check")

	// Auto-shutdown after 30s for demo purposes
	server := &http.Server{Handler: mux}
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("\nAuto-shutting down (demo timeout)...")
		server.Close()
	}()

	if err := server.Serve(listener); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
