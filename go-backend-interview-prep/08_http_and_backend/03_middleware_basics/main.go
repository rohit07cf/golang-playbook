package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// --- Middleware: logging ---
func loggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("  [LOG] %s %s  %s\n", r.Method, r.URL.Path, time.Since(start))
	})
}

// --- Middleware: timing header ---
func timingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		w.Header().Set("X-Response-Time", time.Since(start).String())
	})
}

// --- Middleware: panic recovery ---
func recoveryMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("  [RECOVER] panic: %v\n", err)
				http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// chain applies middleware in order: first listed = outermost
func chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

// --- Handlers ---

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello from middleware demo!")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/panic", panicHandler)

	// Chain: recovery -> logging -> timing -> mux
	handler := chain(mux, recoveryMW, loggingMW, timingMW)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	addr := ln.Addr().String()
	fmt.Printf("Server on http://%s\n", addr)
	fmt.Println("  GET /hello  -> greeting (logged + timed)")
	fmt.Println("  GET /panic  -> triggers recovery middleware")

	server := &http.Server{Handler: handler}
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("\nAuto-shutdown (demo)...")
		server.Close()
	}()

	// Demo: self-test
	go func() {
		time.Sleep(100 * time.Millisecond)
		base := "http://" + addr
		fmt.Println("\n--- Demo requests ---")

		resp, _ := http.Get(base + "/hello")
		if resp != nil {
			resp.Body.Close()
			fmt.Println("  /hello -> status", resp.StatusCode)
		}

		resp2, _ := http.Get(base + "/panic")
		if resp2 != nil {
			resp2.Body.Close()
			fmt.Println("  /panic -> status", resp2.StatusCode, "(recovered)")
		}

		fmt.Println("\nDemo done. Server still running (30s timeout)...")
	}()

	if err := server.Serve(ln); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
