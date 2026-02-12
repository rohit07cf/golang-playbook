package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// apiKey would come from env in production
const apiKey = "secret123"

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// apiKeyAuth middleware checks X-API-Key header.
func apiKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")

		// Constant-time comparison prevents timing attacks
		if subtle.ConstantTimeCompare([]byte(key), []byte(apiKey)) != 1 {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "unauthorized: invalid or missing API key",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "welcome, authenticated user!",
		"secret":  "the answer is 42",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func main() {
	// Protected routes (behind auth)
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/protected", protectedHandler)

	// Public routes (no auth)
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/health", healthHandler)
	mainMux.Handle("/protected", apiKeyAuth(protectedMux))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /health     -> public (no auth)")
	fmt.Println("  GET /protected  -> requires X-API-Key header")
	fmt.Printf("  API key: %s\n", apiKey)

	server := &http.Server{Handler: mainMux}

	// Self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Public endpoint
		resp, _ := http.Get(base + "/health")
		if resp != nil {
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("  /health (no key)  -> %d: %v\n", resp.StatusCode, body)
		}

		// Protected: no key
		resp2, _ := http.Get(base + "/protected")
		if resp2 != nil {
			var body map[string]string
			json.NewDecoder(resp2.Body).Decode(&body)
			resp2.Body.Close()
			fmt.Printf("  /protected (no key) -> %d: %v\n", resp2.StatusCode, body)
		}

		// Protected: wrong key
		req, _ := http.NewRequest("GET", base+"/protected", nil)
		req.Header.Set("X-API-Key", "wrong")
		resp3, _ := http.DefaultClient.Do(req)
		if resp3 != nil {
			var body map[string]string
			json.NewDecoder(resp3.Body).Decode(&body)
			resp3.Body.Close()
			fmt.Printf("  /protected (wrong)  -> %d: %v\n", resp3.StatusCode, body)
		}

		// Protected: correct key
		req2, _ := http.NewRequest("GET", base+"/protected", nil)
		req2.Header.Set("X-API-Key", apiKey)
		resp4, _ := http.DefaultClient.Do(req2)
		if resp4 != nil {
			var body map[string]string
			json.NewDecoder(resp4.Body).Decode(&body)
			resp4.Body.Close()
			fmt.Printf("  /protected (valid)  -> %d: %v\n", resp4.StatusCode, body)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
