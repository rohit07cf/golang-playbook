package main

import (
	"fmt"
	"sync"
)

// --- Example 1: safe counter with Mutex ---

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// --- Example 2: read-heavy cache with RWMutex ---

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // shared read lock
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock() // exclusive write lock
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {
	// --- Safe counter ---
	fmt.Println("=== Safe Counter (Mutex) ===")
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("  counter:", counter.Value()) // always 1000

	// --- RWMutex cache ---
	fmt.Println("\n=== Cache (RWMutex) ===")
	cache := NewCache()
	cache.Set("name", "alice")
	cache.Set("role", "engineer")

	// Multiple readers can run concurrently
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			if val, ok := cache.Get("name"); ok {
				fmt.Printf("  reader %d: name=%s\n", id, val)
			}
		}(i)
	}
	wg2.Wait()

	// Write blocks all readers
	cache.Set("name", "bob")
	val, _ := cache.Get("name")
	fmt.Println("  after write: name =", val)
}
