package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Item is the resource we CRUD.
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Store holds items in memory, protected by a mutex.
var (
	store  = map[int]Item{}
	mu     sync.RWMutex
	nextID = 1
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// --- Handlers ---

func listItems(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	items := make([]Item, 0, len(store))
	for _, item := range store {
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}
	mu.RLock()
	item, ok := store[id]
	mu.RUnlock()
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name required"})
		return
	}

	mu.Lock()
	item := Item{ID: nextID, Name: req.Name, Price: req.Price}
	store[nextID] = item
	nextID++
	mu.Unlock()

	writeJSON(w, http.StatusCreated, item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	item, ok := store[id]
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Price != 0 {
		item.Price = req.Price
	}
	store[id] = item
	writeJSON(w, http.StatusOK, item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if _, ok := store[id]; !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	delete(store, id)
	writeJSON(w, http.StatusOK, map[string]string{"deleted": strconv.Itoa(id)})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", listItems)
	mux.HandleFunc("GET /items/{id}", getItem)
	mux.HandleFunc("POST /items", createItem)
	mux.HandleFunc("PUT /items/{id}", updateItem)
	mux.HandleFunc("DELETE /items/{id}", deleteItem)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("CRUD API on http://%s\n", addr)
	fmt.Println("  POST   /items       -> create")
	fmt.Println("  GET    /items       -> list")
	fmt.Println("  GET    /items/{id}  -> get")
	fmt.Println("  PUT    /items/{id}  -> update")
	fmt.Println("  DELETE /items/{id}  -> delete")

	server := &http.Server{Handler: mux}

	// Self-test demo
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
			fmt.Printf("  CREATE: %+v (status %d)\n", item, resp.StatusCode)
		}

		// Create another
		resp2, _ := http.Post(base+"/items", "application/json",
			strings.NewReader(`{"name":"gadget","price":19.99}`))
		if resp2 != nil {
			resp2.Body.Close()
		}

		// List
		resp3, _ := http.Get(base + "/items")
		if resp3 != nil {
			var items []Item
			json.NewDecoder(resp3.Body).Decode(&items)
			resp3.Body.Close()
			fmt.Printf("  LIST:   %d items\n", len(items))
		}

		// Get
		resp4, _ := http.Get(base + "/items/1")
		if resp4 != nil {
			var item Item
			json.NewDecoder(resp4.Body).Decode(&item)
			resp4.Body.Close()
			fmt.Printf("  GET:    %+v\n", item)
		}

		// Update
		req, _ := http.NewRequest("PUT", base+"/items/1",
			strings.NewReader(`{"name":"super widget","price":14.99}`))
		req.Header.Set("Content-Type", "application/json")
		resp5, _ := http.DefaultClient.Do(req)
		if resp5 != nil {
			var item Item
			json.NewDecoder(resp5.Body).Decode(&item)
			resp5.Body.Close()
			fmt.Printf("  UPDATE: %+v\n", item)
		}

		// Delete
		delReq, _ := http.NewRequest("DELETE", base+"/items/2", nil)
		resp6, _ := http.DefaultClient.Do(delReq)
		if resp6 != nil {
			var result map[string]string
			json.NewDecoder(resp6.Body).Decode(&result)
			resp6.Body.Close()
			fmt.Printf("  DELETE: %v\n", result)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
