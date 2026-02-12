package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// --- Cache Entry ---

type entry struct {
	key       string
	value     interface{}
	expiresAt time.Time
	element   *list.Element // pointer into LRU list
}

// --- Cache ---

type Cache struct {
	mu       sync.Mutex
	items    map[string]*entry
	order    *list.List // front = most recent, back = least recent
	capacity int
}

func NewCache(capacity int) *Cache {
	c := &Cache{
		items:    make(map[string]*entry),
		order:    list.New(),
		capacity: capacity,
	}
	// Periodic cleanup of expired entries
	go c.cleanupLoop(1 * time.Second)
	return c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Update existing entry
	if e, ok := c.items[key]; ok {
		e.value = value
		e.expiresAt = time.Now().Add(ttl)
		c.order.MoveToFront(e.element)
		return
	}

	// Evict if at capacity
	if len(c.items) >= c.capacity {
		c.evictLRU()
	}

	// Add new entry
	e := &entry{
		key:       key,
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	e.element = c.order.PushFront(e)
	c.items[key] = e
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.items[key]
	if !ok {
		return nil, false
	}

	// Lazy expiration check
	if time.Now().After(e.expiresAt) {
		c.removeEntry(e)
		return nil, false
	}

	// Move to front (most recently used)
	c.order.MoveToFront(e.element)
	return e.value, true
}

func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

// evictLRU removes the least recently used entry. Caller must hold lock.
func (c *Cache) evictLRU() {
	back := c.order.Back()
	if back == nil {
		return
	}
	e := back.Value.(*entry)
	fmt.Printf("  [evict] key=%q (LRU)\n", e.key)
	c.removeEntry(e)
}

// removeEntry deletes an entry from both map and list. Caller must hold lock.
func (c *Cache) removeEntry(e *entry) {
	c.order.Remove(e.element)
	delete(c.items, e.key)
}

// cleanupLoop periodically removes expired entries.
func (c *Cache) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, e := range c.items {
			if now.After(e.expiresAt) {
				fmt.Printf("  [cleanup] expired key=%q\n", key)
				c.removeEntry(e)
			}
		}
		c.mu.Unlock()
	}
}

// --- Demo ---

func main() {
	cache := NewCache(3) // max 3 entries

	fmt.Println("=== cache demo (capacity=3) ===")

	// Set entries with different TTLs
	cache.Set("user:1", "Alice", 5*time.Second)
	cache.Set("user:2", "Bob", 2*time.Second)
	cache.Set("user:3", "Charlie", 5*time.Second)
	fmt.Printf("set 3 entries, cache size: %d\n\n", cache.Len())

	// Cache hit
	if v, ok := cache.Get("user:1"); ok {
		fmt.Printf("GET user:1 -> %v (hit)\n", v)
	}
	if v, ok := cache.Get("user:2"); ok {
		fmt.Printf("GET user:2 -> %v (hit)\n", v)
	}

	// Cache miss
	if _, ok := cache.Get("user:99"); !ok {
		fmt.Printf("GET user:99 -> (miss)\n")
	}

	// Trigger LRU eviction (add 4th entry, capacity is 3)
	fmt.Println("\nadding user:4 (should evict LRU)...")
	cache.Set("user:4", "Diana", 5*time.Second)
	fmt.Printf("cache size: %d\n", cache.Len())

	// user:3 should be evicted (least recently accessed)
	if _, ok := cache.Get("user:3"); !ok {
		fmt.Printf("GET user:3 -> (miss, was evicted as LRU)\n")
	}

	// Wait for TTL expiration
	fmt.Println("\nwaiting 3s for user:2 TTL to expire...")
	time.Sleep(3 * time.Second)

	if _, ok := cache.Get("user:2"); !ok {
		fmt.Printf("GET user:2 -> (miss, expired)\n")
	}

	fmt.Printf("\nfinal cache size: %d\n", cache.Len())
	fmt.Println("\ndemo done")
}
