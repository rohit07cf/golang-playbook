package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Limit body size (1 MB)
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Check Content-Type
	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Content-Type must be application/json",
		})
		return
	}

	// Decode JSON
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid JSON: " + err.Error(),
		})
		return
	}

	// Validate fields
	errors := []string{}
	if req.Name == "" {
		errors = append(errors, "name is required")
	}
	if req.Email == "" {
		errors = append(errors, "email is required")
	} else if !strings.Contains(req.Email, "@") {
		errors = append(errors, "email must contain @")
	}
	if len(errors) > 0 {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"error":  "validation failed",
			"fields": errors,
		})
		return
	}

	// Success
	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "user created",
		"name":    req.Name,
		"email":   req.Email,
	})
}

// queryHandler shows query param parsing
func queryHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page := q.Get("page")
	limit := q.Get("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"page":  page,
		"limit": limit,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /search", queryHandler)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  POST /users   -> create user (JSON body)")
	fmt.Println("  GET  /search  -> query params demo")

	server := &http.Server{Handler: mux}

	// Self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo ---")

		// Valid request
		body := `{"name":"Alice","email":"alice@example.com"}`
		resp, _ := http.Post(base+"/users", "application/json", strings.NewReader(body))
		if resp != nil {
			var result map[string]any
			json.NewDecoder(resp.Body).Decode(&result)
			resp.Body.Close()
			fmt.Printf("  valid:   %v\n", result)
		}

		// Missing fields
		resp2, _ := http.Post(base+"/users", "application/json", strings.NewReader(`{}`))
		if resp2 != nil {
			var result map[string]any
			json.NewDecoder(resp2.Body).Decode(&result)
			resp2.Body.Close()
			fmt.Printf("  invalid: %v\n", result)
		}

		// Bad email
		resp3, _ := http.Post(base+"/users", "application/json", strings.NewReader(`{"name":"Bob","email":"nope"}`))
		if resp3 != nil {
			var result map[string]any
			json.NewDecoder(resp3.Body).Decode(&result)
			resp3.Body.Close()
			fmt.Printf("  bad email: %v\n", result)
		}

		// Query params
		resp4, _ := http.Get(base + "/search?page=2&limit=25")
		if resp4 != nil {
			var result map[string]any
			json.NewDecoder(resp4.Body).Decode(&result)
			resp4.Body.Close()
			fmt.Printf("  query:   %v\n", result)
		}

		fmt.Println("\nDemo done.")
		server.Close()
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
