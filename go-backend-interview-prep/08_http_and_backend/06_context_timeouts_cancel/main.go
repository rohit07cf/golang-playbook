package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// simulateWork pretends to do something that respects context cancellation.
func simulateWork(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "done", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// fastHandler completes within the timeout.
func fastHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	result, err := simulateWork(ctx, 100*time.Millisecond)
	if err != nil {
		writeJSON(w, http.StatusGatewayTimeout, map[string]string{"error": "timeout"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"result": result, "endpoint": "fast"})
}

// slowHandler exceeds the timeout -> context cancelled.
func slowHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	result, err := simulateWork(ctx, 5*time.Second) // too slow!
	if err != nil {
		writeJSON(w, http.StatusGatewayTimeout, map[string]string{
			"error": "request timed out",
			"cause": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"result": result})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /fast", fastHandler)
	mux.HandleFunc("GET /slow", slowHandler)

	// Server-level timeouts (defense in depth)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /fast -> completes within timeout")
	fmt.Println("  GET /slow -> exceeds timeout (504)")

	// Self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Fast: should succeed
		resp, _ := http.Get(base + "/fast")
		if resp != nil {
			var result map[string]string
			json.NewDecoder(resp.Body).Decode(&result)
			resp.Body.Close()
			fmt.Printf("  /fast -> %d: %v\n", resp.StatusCode, result)
		}

		// Slow: should timeout
		resp2, _ := http.Get(base + "/slow")
		if resp2 != nil {
			var result map[string]string
			json.NewDecoder(resp2.Body).Decode(&result)
			resp2.Body.Close()
			fmt.Printf("  /slow -> %d: %v\n", resp2.StatusCode, result)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
