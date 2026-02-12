package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// --- Request ID ---

type ctxKey string

const reqIDKey ctxKey = "request_id"

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// --- Middleware: Auth ---

var validAPIKeys = map[string]bool{
	"key-abc-123": true,
	"key-xyz-789": true,
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if !validAPIKeys[key] {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "invalid or missing API key",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --- Middleware: Rate Limit (simple fixed window) ---

type rateLimiter struct {
	mu       sync.Mutex
	counts   map[string]int
	limit    int
	window   time.Duration
	lastReset time.Time
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		counts:    make(map[string]int),
		limit:     limit,
		window:    window,
		lastReset: time.Now(),
	}
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Since(rl.lastReset) > rl.window {
		rl.counts = make(map[string]int)
		rl.lastReset = time.Now()
	}

	rl.counts[key]++
	return rl.counts[key] <= rl.limit
}

func rateLimitMiddleware(rl *rateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr
			if !rl.allow(key) {
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "rate limit exceeded",
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// --- Middleware: Request ID ---

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := generateID()
		ctx := context.WithValue(r.Context(), reqIDKey, id)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// --- Middleware: Timeout ---

func timeoutMiddleware(d time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// --- Downstream Handlers (toy services) ---

func userHandler(w http.ResponseWriter, r *http.Request) {
	reqID, _ := r.Context().Value(reqIDKey).(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"service":    "users",
		"path":       r.URL.Path,
		"request_id": reqID,
		"data":       "user list placeholder",
	})
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	reqID, _ := r.Context().Value(reqIDKey).(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"service":    "orders",
		"path":       r.URL.Path,
		"request_id": reqID,
		"data":       "order list placeholder",
	})
}

// --- Router ---

func gatewayRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/users"):
			userHandler(w, r)
		case strings.HasPrefix(r.URL.Path, "/api/orders"):
			orderHandler(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "route not found",
			})
		}
	}
}

// --- Chain helper ---

func chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// --- Demo ---

func main() {
	rl := newRateLimiter(5, 10*time.Second) // 5 req per 10s window

	gateway := chain(
		gatewayRouter(),
		authMiddleware,
		rateLimitMiddleware(rl),
		requestIDMiddleware,
		timeoutMiddleware(2*time.Second),
	)

	mux := http.NewServeMux()
	mux.Handle("/", gateway)

	go func() {
		time.Sleep(500 * time.Millisecond)
		client := &http.Client{Timeout: 3 * time.Second}

		fmt.Print("--- API Gateway Demo ---\n\n")

		// 1. No API key
		req, _ := http.NewRequest("GET", "http://localhost:9007/api/users", nil)
		resp, _ := client.Do(req)
		fmt.Printf("GET /api/users (no key): status=%d\n", resp.StatusCode)
		resp.Body.Close()

		// 2. Valid API key -- users
		req, _ = http.NewRequest("GET", "http://localhost:9007/api/users", nil)
		req.Header.Set("X-API-Key", "key-abc-123")
		resp, _ = client.Do(req)
		var body map[string]string
		json.NewDecoder(resp.Body).Decode(&body)
		resp.Body.Close()
		fmt.Printf("GET /api/users (valid key): status=%d service=%s req_id=%s\n",
			resp.StatusCode, body["service"], body["request_id"])

		// 3. Valid API key -- orders
		req, _ = http.NewRequest("GET", "http://localhost:9007/api/orders", nil)
		req.Header.Set("X-API-Key", "key-abc-123")
		resp, _ = client.Do(req)
		json.NewDecoder(resp.Body).Decode(&body)
		resp.Body.Close()
		fmt.Printf("GET /api/orders (valid key): status=%d service=%s\n",
			resp.StatusCode, body["service"])

		// 4. Unknown route
		req, _ = http.NewRequest("GET", "http://localhost:9007/api/unknown", nil)
		req.Header.Set("X-API-Key", "key-abc-123")
		resp, _ = client.Do(req)
		fmt.Printf("GET /api/unknown: status=%d\n", resp.StatusCode)
		resp.Body.Close()

		// 5. Rate limit test
		fmt.Println("\n--- Rate limit test (limit=5/10s) ---")
		for i := 1; i <= 7; i++ {
			req, _ = http.NewRequest("GET", "http://localhost:9007/api/users", nil)
			req.Header.Set("X-API-Key", "key-xyz-789")
			resp, _ = client.Do(req)
			fmt.Printf("req %d: status=%d\n", i, resp.StatusCode)
			resp.Body.Close()
		}

		log.Println("demo done")
		go func() { time.Sleep(100 * time.Millisecond); log.Fatal("exit") }()
	}()

	fmt.Println("api gateway on :9007")
	log.Fatal(http.ListenAndServe(":9007", mux))
}
