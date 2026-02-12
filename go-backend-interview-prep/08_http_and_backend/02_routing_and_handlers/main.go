package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// In-memory store (mutex-protected)
var (
	items   = map[string]string{"1": "alpha", "2": "bravo"}
	mu      sync.RWMutex
	nextID  = 3
)

// --- Handlers ---

func listItems(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	writeJSON(w, http.StatusOK, items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	mu.RLock()
	defer mu.RUnlock()
	val, ok := items[id]
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"id": id, "name": val})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var body struct{ Name string }
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name required"})
		return
	}
	mu.Lock()
	id := fmt.Sprintf("%d", nextID)
	nextID++
	items[id] = body.Name
	mu.Unlock()
	writeJSON(w, http.StatusCreated, map[string]string{"id": id, "name": body.Name})
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	mu.Lock()
	defer mu.Unlock()
	if _, ok := items[id]; !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	delete(items, id)
	writeJSON(w, http.StatusOK, map[string]string{"deleted": id})
}

// Catch-all: returns 404 JSON for unmatched routes
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotFound, map[string]string{
		"error": "route not found",
		"path":  r.URL.Path,
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// methodCheck wraps a handler to restrict to specific methods (pre-1.22 style)
func methodCheck(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(r.Method, method) {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error":   "method not allowed",
				"allowed": method,
			})
			return
		}
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	// Go 1.22+ style: "METHOD /path"
	mux.HandleFunc("GET /items", listItems)
	mux.HandleFunc("GET /items/{id}", getItem)
	mux.HandleFunc("POST /items", createItem)
	mux.HandleFunc("DELETE /items/{id}", deleteItem)

	// Catch-all for unknown routes
	mux.HandleFunc("/", notFoundHandler)

	// Also show old-style method check for interview reference
	_ = methodCheck // referenced above for pre-1.22 fallback

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET    /items       -> list all")
	fmt.Println("  GET    /items/{id}  -> get one")
	fmt.Println("  POST   /items       -> create")
	fmt.Println("  DELETE /items/{id}  -> delete")

	server := &http.Server{Handler: mux}
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("\nAuto-shutdown (demo)...")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
