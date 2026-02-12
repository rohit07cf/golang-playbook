package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// slowHandler simulates a request that takes time to complete.
func slowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("  [handler] slow request started, sleeping 2s...")
	time.Sleep(2 * time.Second)
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "slow request completed",
	})
	fmt.Println("  [handler] slow request done")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/slow", slowHandler)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()

	server := &http.Server{Handler: mux}

	// --- Graceful shutdown on signal ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-quit
		fmt.Printf("\n  [shutdown] received signal: %s\n", sig)
		fmt.Println("  [shutdown] draining in-flight requests (10s timeout)...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("  [shutdown] error: %v\n", err)
		} else {
			fmt.Println("  [shutdown] graceful shutdown complete")
		}
	}()

	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /health -> health check")
	fmt.Println("  GET /slow   -> 2s slow request")
	fmt.Println("  Press Ctrl+C to trigger graceful shutdown")

	// Self-test: start a slow request, then trigger shutdown mid-flight
	go func() {
		time.Sleep(200 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Start slow request in background
		done := make(chan struct{})
		go func() {
			defer close(done)
			resp, err := http.Get(base + "/slow")
			if err != nil {
				fmt.Printf("  slow request error: %v\n", err)
				return
			}
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("  slow request result: %v (status %d)\n", body, resp.StatusCode)
		}()

		// Wait a bit then signal shutdown while slow request is in-flight
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  [demo] sending SIGINT while slow request is in-flight...")
		quit <- syscall.SIGINT

		// Wait for slow request to complete (should finish under graceful drain)
		<-done
	}()

	// Serve blocks until Shutdown is called
	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait a moment for shutdown goroutine to finish logging
	time.Sleep(100 * time.Millisecond)
	fmt.Println("\nServer exited.")
}
