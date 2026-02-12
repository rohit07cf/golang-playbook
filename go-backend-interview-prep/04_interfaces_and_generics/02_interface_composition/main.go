package main

import "fmt"

// --- Small, focused interfaces ---
type Saver interface {
	Save(data string) error
}

type Loader interface {
	Load(id string) (string, error)
}

type Deleter interface {
	Delete(id string) error
}

// --- Composed interface: embeds all three ---
type Storage interface {
	Saver
	Loader
	Deleter
}

// --- Concrete type satisfying Storage ---
type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]string)}
}

func (m *MemoryStore) Save(data string) error {
	id := fmt.Sprintf("key-%d", len(m.data)+1)
	m.data[id] = data
	fmt.Printf("  saved: %s -> %q\n", id, data)
	return nil
}

func (m *MemoryStore) Load(id string) (string, error) {
	val, ok := m.data[id]
	if !ok {
		return "", fmt.Errorf("not found: %s", id)
	}
	return val, nil
}

func (m *MemoryStore) Delete(id string) error {
	delete(m.data, id)
	fmt.Printf("  deleted: %s\n", id)
	return nil
}

// --- Functions accepting sub-interfaces ---
func saveRecord(s Saver, data string) {
	s.Save(data)
}

func loadRecord(l Loader, id string) {
	val, err := l.Load(id)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	fmt.Printf("  loaded: %s -> %q\n", id, val)
}

func main() {
	store := NewMemoryStore()

	// --- Use as composed interface ---
	fmt.Println("--- Full Storage interface ---")
	var s Storage = store
	s.Save("hello")
	s.Save("world")

	// --- Use as sub-interface ---
	fmt.Println("\n--- As Loader only ---")
	loadRecord(store, "key-1")
	loadRecord(store, "key-99")

	// --- Use as Saver only ---
	fmt.Println("\n--- As Saver only ---")
	saveRecord(store, "another")

	// --- Use as Deleter ---
	fmt.Println("\n--- As Deleter ---")
	store.Delete("key-1")

	// --- Pass where any sub-interface is expected ---
	// MemoryStore satisfies Saver, Loader, Deleter, AND Storage.
	fmt.Println("\n--- Satisfies all sub-interfaces ---")
	var saver Saver = store
	var loader Loader = store
	_ = saver
	_ = loader
	fmt.Println("MemoryStore satisfies Saver, Loader, Deleter, and Storage")
}
