package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// --- Store Interface ---

type URLStore interface {
	Save(code, url string)
	Load(code string) (string, bool)
	Exists(code string) bool
}

// --- In-Memory Store ---

type MemoryStore struct {
	mu   sync.RWMutex
	urls map[string]string // code -> url
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{urls: make(map[string]string)}
}

func (s *MemoryStore) Save(code, url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[code] = url
}

func (s *MemoryStore) Load(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.urls[code]
	return url, ok
}

func (s *MemoryStore) Exists(code string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.urls[code]
	return ok
}

// --- Code Generator ---

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[n.Int64()]
	}
	return string(b)
}

func generateUniqueCode(store URLStore, length int) string {
	for i := 0; i < 10; i++ { // max 10 retries for collision
		code := generateCode(length)
		if !store.Exists(code) {
			return code
		}
	}
	// Fallback: longer code
	return generateCode(length + 2)
}

// --- Handlers ---

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

func shortenHandler(store URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
			http.Error(w, `{"error":"provide a valid url"}`, http.StatusBadRequest)
			return
		}

		code := generateUniqueCode(store, 6)
		store.Save(code, req.URL)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ShortenResponse{
			Code:     code,
			ShortURL: "http://localhost:9002/" + code,
		})
		log.Printf("shortened: %s -> %s", req.URL, code)
	}
}

func resolveHandler(store URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.URL.Path, "/r/")
		if code == "" {
			http.Error(w, `{"error":"code required"}`, http.StatusBadRequest)
			return
		}

		url, ok := store.Load(code)
		if !ok {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"code":         code,
			"original_url": url,
		})
		log.Printf("resolved: %s -> %s", code, url)
	}
}

// --- Demo ---

func main() {
	store := NewMemoryStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", shortenHandler(store))
	mux.HandleFunc("/r/", resolveHandler(store))

	go func() {
		time.Sleep(500 * time.Millisecond)
		client := &http.Client{Timeout: 2 * time.Second}

		// Shorten a few URLs
		urls := []string{
			"https://go.dev/doc/effective_go",
			"https://pkg.go.dev/net/http",
			"https://example.com/very/long/path/to/resource",
		}

		codes := []string{}
		for _, u := range urls {
			body := fmt.Sprintf(`{"url":"%s"}`, u)
			resp, err := client.Post(
				"http://localhost:9002/shorten",
				"application/json",
				strings.NewReader(body),
			)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			}
			var result ShortenResponse
			json.NewDecoder(resp.Body).Decode(&result)
			resp.Body.Close()
			fmt.Printf("POST /shorten %s -> code=%s\n", u, result.Code)
			codes = append(codes, result.Code)
		}

		// Resolve them
		fmt.Println()
		for _, code := range codes {
			resp, err := client.Get("http://localhost:9002/r/" + code)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			}
			var result map[string]string
			json.NewDecoder(resp.Body).Decode(&result)
			resp.Body.Close()
			fmt.Printf("GET /r/%s -> %s\n", code, result["original_url"])
		}

		// Try a missing code
		fmt.Println()
		resp, _ := client.Get("http://localhost:9002/r/missing")
		fmt.Printf("GET /r/missing -> status=%d\n", resp.StatusCode)
		resp.Body.Close()

		log.Println("demo done")
		go func() { time.Sleep(100 * time.Millisecond); log.Fatal("exit") }()
	}()

	fmt.Println("url shortener on :9002")
	log.Fatal(http.ListenAndServe(":9002", mux))
}
