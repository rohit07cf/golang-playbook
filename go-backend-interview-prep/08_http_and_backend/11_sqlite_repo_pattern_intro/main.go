package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// --- Domain ---

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// --- Repository interface ---
// In production, you'd have SQLiteRepo, PostgresRepo, etc.
// All implement this interface -- handlers never know which one.

type ItemRepo interface {
	Create(name string, price float64) (Item, error)
	GetByID(id int) (Item, error)
	List() ([]Item, error)
	Delete(id int) error
}

// --- In-memory implementation ---
// Go's stdlib has no SQLite driver. Real SQLite needs an external package.
// We demonstrate the pattern with a map.

type MemoryRepo struct {
	mu     sync.RWMutex
	items  map[int]Item
	nextID int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{items: map[int]Item{}, nextID: 1}
}

var ErrNotFound = errors.New("item not found")

func (r *MemoryRepo) Create(name string, price float64) (Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	item := Item{ID: r.nextID, Name: name, Price: price}
	r.items[r.nextID] = item
	r.nextID++
	return item, nil
}

func (r *MemoryRepo) GetByID(id int) (Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[id]
	if !ok {
		return Item{}, ErrNotFound
	}
	return item, nil
}

func (r *MemoryRepo) List() ([]Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]Item, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	return items, nil
}

func (r *MemoryRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}

// --- Service layer (business logic) ---
// Thin in this demo, but in production this is where validation,
// business rules, and cross-cutting concerns live.

type ItemService struct {
	repo ItemRepo
}

func (s *ItemService) CreateItem(name string, price float64) (Item, error) {
	if name == "" {
		return Item{}, errors.New("name is required")
	}
	if price < 0 {
		return Item{}, errors.New("price must be non-negative")
	}
	return s.repo.Create(name, price)
}

func (s *ItemService) GetItem(id int) (Item, error) {
	return s.repo.GetByID(id)
}

func (s *ItemService) ListItems() ([]Item, error) {
	return s.repo.List()
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repo.Delete(id)
}

// --- HTTP Handlers ---

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func makeHandlers(svc *ItemService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		items, err := svc.ListItems()
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, items)
	})

	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PathValue("id"))
		item, err := svc.GetItem(id)
		if errors.Is(err, ErrNotFound) {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, item)
	})

	mux.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid json"})
			return
		}
		item, err := svc.CreateItem(req.Name, req.Price)
		if err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, item)
	})

	mux.HandleFunc("DELETE /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PathValue("id"))
		err := svc.DeleteItem(id)
		if errors.Is(err, ErrNotFound) {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, map[string]string{"deleted": strconv.Itoa(id)})
	})

	return mux
}

func main() {
	// Wire up: repo -> service -> handlers
	repo := NewMemoryRepo()
	svc := &ItemService{repo: repo}
	handler := makeHandlers(svc)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Repo Pattern API on http://%s\n", addr)
	fmt.Println("  Architecture: handler -> service -> repository (interface)")
	fmt.Println("  Storage: in-memory map (swap to SQLite/Postgres by changing repo)")

	server := &http.Server{Handler: handler}

	// Self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Create
		resp, _ := http.Post(base+"/items", "application/json",
			strings.NewReader(`{"name":"widget","price":9.99}`))
		if resp != nil {
			var item Item
			json.NewDecoder(resp.Body).Decode(&item)
			resp.Body.Close()
			fmt.Printf("  CREATE: %+v\n", item)
		}

		http.Post(base+"/items", "application/json",
			strings.NewReader(`{"name":"gadget","price":19.99}`))

		// List
		resp2, _ := http.Get(base + "/items")
		if resp2 != nil {
			var items []Item
			json.NewDecoder(resp2.Body).Decode(&items)
			resp2.Body.Close()
			fmt.Printf("  LIST:   %d items\n", len(items))
		}

		// Get
		resp3, _ := http.Get(base + "/items/1")
		if resp3 != nil {
			var item Item
			json.NewDecoder(resp3.Body).Decode(&item)
			resp3.Body.Close()
			fmt.Printf("  GET:    %+v\n", item)
		}

		// Delete
		delReq, _ := http.NewRequest("DELETE", base+"/items/2", nil)
		resp4, _ := http.DefaultClient.Do(delReq)
		if resp4 != nil {
			var result map[string]string
			json.NewDecoder(resp4.Body).Decode(&result)
			resp4.Body.Close()
			fmt.Printf("  DELETE: %v\n", result)
		}

		// 404
		resp5, _ := http.Get(base + "/items/99")
		if resp5 != nil {
			resp5.Body.Close()
			fmt.Printf("  GET 99: status %d\n", resp5.StatusCode)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
