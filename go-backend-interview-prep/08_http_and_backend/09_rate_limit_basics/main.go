package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// TokenBucket implements a simple rate limiter.
type TokenBucket struct {
	mu       sync.Mutex
	tokens   int
	max      int
	interval time.Duration
}

func NewTokenBucket(max int, refillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		tokens:   max,
		max:      max,
		interval: refillInterval,
	}
	// Refill tokens periodically
	go func() {
		ticker := time.NewTicker(refillInterval)
		defer ticker.Stop()
		for range ticker.C {
			tb.mu.Lock()
			tb.tokens = tb.max
			tb.mu.Unlock()
		}
	}()
	return tb
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) Remaining() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// rateLimitMiddleware wraps a handler with token bucket checking.
func rateLimitMiddleware(limiter *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.Header().Set("Retry-After", fmt.Sprintf("%.0f", limiter.interval.Seconds()))
			writeJSON(w, http.StatusTooManyRequests, map[string]string{
				"error":       "rate limit exceeded",
				"retry_after": fmt.Sprintf("%.0fs", limiter.interval.Seconds()),
			})
			return
		}
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Remaining()))
		next.ServeHTTP(w, r)
	})
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "API response",
	})
}

func main() {
	// 5 requests per 3 seconds
	limiter := NewTokenBucket(5, 3*time.Second)

	mux := http.NewServeMux()
	mux.HandleFunc("/api", apiHandler)

	handler := rateLimitMiddleware(limiter, mux)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /api -> rate limited (5 req / 3s)")

	server := &http.Server{Handler: handler}

	// Self-test: send 8 requests rapidly
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo: 8 rapid requests ---")

		for i := 1; i <= 8; i++ {
			resp, err := http.Get(base + "/api")
			if err != nil {
				fmt.Printf("  request %d: error: %v\n", i, err)
				continue
			}
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			remaining := resp.Header.Get("X-RateLimit-Remaining")
			fmt.Printf("  request %d: status=%d remaining=%s msg=%s\n",
				i, resp.StatusCode, remaining, body["message"]+body["error"])
		}

		// Wait for refill
		fmt.Println("\n  waiting 3s for token refill...")
		time.Sleep(3 * time.Second)

		resp, _ := http.Get(base + "/api")
		if resp != nil {
			resp.Body.Close()
			fmt.Printf("  after refill: status=%d\n", resp.StatusCode)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
