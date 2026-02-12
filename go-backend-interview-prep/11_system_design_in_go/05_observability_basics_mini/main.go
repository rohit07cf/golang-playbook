package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// --- Request ID ---

type ctxKey string

const requestIDKey ctxKey = "request_id"

func generateRequestID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func getRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return "unknown"
}

// --- Metrics ---

type Metrics struct {
	totalRequests atomic.Int64
	totalErrors   atomic.Int64
	totalLatencyMs atomic.Int64
}

var metrics Metrics

// --- Middleware: Request ID ---

func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := generateRequestID()
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		w.Header().Set("X-Request-ID", id)
		next(w, r.WithContext(ctx))
	}
}

// --- Middleware: Logging + Metrics ---

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, statusCode: 200}

		next(recorder, r)

		latency := time.Since(start)
		reqID := getRequestID(r.Context())

		// Update metrics
		metrics.totalRequests.Add(1)
		metrics.totalLatencyMs.Add(latency.Milliseconds())
		if recorder.statusCode >= 400 {
			metrics.totalErrors.Add(1)
		}

		// Structured log line
		log.Printf("request_id=%s method=%s path=%s status=%d latency=%v",
			reqID, r.Method, r.URL.Path, recorder.statusCode, latency)
	}
}

// --- Handlers ---

func helloHandler(w http.ResponseWriter, r *http.Request) {
	reqID := getRequestID(r.Context())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":    "hello",
		"request_id": reqID,
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "something went wrong",
	})
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	total := metrics.totalRequests.Load()
	errors := metrics.totalErrors.Load()
	latency := metrics.totalLatencyMs.Load()

	avgLatency := float64(0)
	if total > 0 {
		avgLatency = float64(latency) / float64(total)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "requests_total %d\n", total)
	fmt.Fprintf(w, "errors_total %d\n", errors)
	fmt.Fprintf(w, "latency_total_ms %d\n", latency)
	fmt.Fprintf(w, "latency_avg_ms %.2f\n", avgLatency)
}

// --- Chain helper ---

func chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// --- Demo ---

func main() {
	mux := http.NewServeMux()

	// Apply middleware chain
	mux.HandleFunc("/hello", chain(helloHandler, requestIDMiddleware, loggingMiddleware))
	mux.HandleFunc("/error", chain(errorHandler, requestIDMiddleware, loggingMiddleware))
	mux.HandleFunc("/metrics", metricsHandler)

	// Demo: simulate requests
	go func() {
		time.Sleep(500 * time.Millisecond)
		client := &http.Client{Timeout: 2 * time.Second}

		fmt.Print("--- sending requests ---\n\n")

		// Normal requests
		for i := 0; i < 5; i++ {
			resp, _ := client.Get("http://localhost:9005/hello")
			var body map[string]string
			json.NewDecoder(resp.Body).Decode(&body)
			resp.Body.Close()
			fmt.Printf("GET /hello -> request_id=%s\n", body["request_id"])
		}

		// Error request
		resp, _ := client.Get("http://localhost:9005/error")
		fmt.Printf("GET /error -> status=%d\n", resp.StatusCode)
		resp.Body.Close()

		// Check metrics
		fmt.Print("\n--- GET /metrics ---\n\n")
		resp, _ = client.Get("http://localhost:9005/metrics")
		buf := make([]byte, 1024)
		n, _ := resp.Body.Read(buf)
		resp.Body.Close()
		fmt.Println(string(buf[:n]))

		log.Println("demo done")
		go func() { time.Sleep(100 * time.Millisecond); log.Fatal("exit") }()
	}()

	fmt.Println("observability demo on :9005")
	log.Fatal(http.ListenAndServe(":9005", mux))
}
