package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// --- Token Bucket ---

// Bucket tracks tokens for a single client.
type Bucket struct {
	tokens     float64
	maxTokens  float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

// Allow checks if a request is allowed and consumes a token.
func (b *Bucket) Allow() bool {
	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.maxTokens {
		b.tokens = b.maxTokens
	}
	b.lastRefill = now

	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

// --- Rate Limiter ---

// RateLimiter manages per-client buckets.
type RateLimiter struct {
	mu         sync.Mutex
	buckets    map[string]*Bucket
	maxTokens  float64
	refillRate float64
}

// NewRateLimiter creates a limiter with given rate and burst.
func NewRateLimiter(refillRate, maxTokens float64) *RateLimiter {
	return &RateLimiter{
		buckets:    make(map[string]*Bucket),
		maxTokens:  maxTokens,
		refillRate: refillRate,
	}
}

// Allow checks if the given client key is within rate limit.
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	b, ok := rl.buckets[key]
	if !ok {
		b = &Bucket{
			tokens:     rl.maxTokens,
			maxTokens:  rl.maxTokens,
			refillRate: rl.refillRate,
			lastRefill: time.Now(),
		}
		rl.buckets[key] = b
	}
	return b.Allow()
}

// --- Middleware ---

func rateLimitMiddleware(rl *RateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.RemoteAddr // per-client by IP
		if !rl.Allow(key) {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "rate limit exceeded",
			})
			return
		}
		next(w, r)
	}
}

// --- Handler ---

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "pong",
	})
}

// --- Demo ---

func main() {
	// 2 tokens/sec, burst of 5
	limiter := NewRateLimiter(2, 5)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", rateLimitMiddleware(limiter, pingHandler))

	// Demo: simulate requests in a goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)
		client := &http.Client{Timeout: 2 * time.Second}
		for i := 1; i <= 8; i++ {
			resp, err := client.Get("http://localhost:9001/ping")
			if err != nil {
				fmt.Printf("req %d: error: %v\n", i, err)
				continue
			}
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("req %d: status=%d body=%v\n", i, resp.StatusCode, body)
		}

		// Wait for refill, then try again
		fmt.Println("\n--- waiting 2s for token refill ---")
		time.Sleep(2 * time.Second)

		resp, err := client.Get("http://localhost:9001/ping")
		if err != nil {
			fmt.Printf("req after wait: error: %v\n", err)
		} else {
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("req after wait: status=%d body=%v\n", resp.StatusCode, body)
		}

		log.Println("demo done -- shutting down")
		go func() { time.Sleep(100 * time.Millisecond); log.Fatal("exit") }()
	}()

	fmt.Println("rate limiter server on :9001 (2 tokens/sec, burst 5)")
	log.Fatal(http.ListenAndServe(":9001", mux))
}
