package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// contextKey is a private type to avoid context key collisions.
type contextKey string

const reqIDKey contextKey = "requestID"

// generateID creates a short random hex string.
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// --- Middleware: request ID ---
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Accept incoming ID or generate new one
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = generateID()
		}
		ctx := context.WithValue(r.Context(), reqIDKey, id)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// statusRecorder captures the status code written by the handler.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

// --- Middleware: structured logging ---
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)

		id, _ := r.Context().Value(reqIDKey).(string)
		fmt.Printf("  [LOG] request_id=%s method=%s path=%s status=%d duration=%s\n",
			id, r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(reqIDKey).(string)
	writeJSON(w, http.StatusOK, map[string]string{
		"message":    "hello",
		"request_id": id,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// Chain: requestID -> logging -> mux
	handler := requestIDMiddleware(loggingMiddleware(mux))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /hello -> greeting with request ID")

	server := &http.Server{Handler: handler}

	// Self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Normal request (server generates ID)
		resp, _ := http.Get(base + "/hello")
		if resp != nil {
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("  response: %v\n", body)
			fmt.Printf("  X-Request-ID header: %s\n", resp.Header.Get("X-Request-ID"))
		}

		// Request with client-provided ID
		req, _ := http.NewRequest("GET", base+"/hello", nil)
		req.Header.Set("X-Request-ID", "client-trace-abc123")
		resp2, _ := http.DefaultClient.Do(req)
		if resp2 != nil {
			var body map[string]string
			json.NewDecoder(resp2.Body).Decode(&body)
			resp2.Body.Close()
			fmt.Printf("  with client ID: %v\n", body)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
