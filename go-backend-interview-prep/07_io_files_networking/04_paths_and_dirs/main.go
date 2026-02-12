package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// --- Example 1: filepath.Join ---
	fmt.Println("=== filepath.Join ===")
	p := filepath.Join("data", "users", "profile.json")
	fmt.Println("  joined:", p)
	fmt.Println("  dir:  ", filepath.Dir(p))
	fmt.Println("  base: ", filepath.Base(p))
	fmt.Println("  ext:  ", filepath.Ext(p))

	// --- Example 2: temp directory + create files ---
	fmt.Println("\n=== Temp directory ===")
	dir, err := os.MkdirTemp("", "paths-demo")
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer os.RemoveAll(dir) // cleanup

	// Create nested structure
	subdir := filepath.Join(dir, "sub", "nested")
	os.MkdirAll(subdir, 0755)

	// Create a few files
	for _, name := range []string{"a.txt", "b.go", "c.json"} {
		path := filepath.Join(dir, name)
		os.WriteFile(path, []byte("content of "+name), 0644)
	}
	os.WriteFile(filepath.Join(subdir, "deep.txt"), []byte("deep file"), 0644)

	fmt.Println("  created temp dir:", dir)

	// --- Example 3: WalkDir ---
	fmt.Println("\n=== WalkDir ===")
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Show relative path for cleaner output
		rel, _ := filepath.Rel(dir, path)
		if d.IsDir() {
			fmt.Printf("  [dir]  %s/\n", rel)
		} else {
			info, _ := d.Info()
			fmt.Printf("  [file] %s (%d bytes)\n", rel, info.Size())
		}
		return nil
	})

	// --- Example 4: glob pattern matching ---
	fmt.Println("\n=== Glob ===")
	matches, _ := filepath.Glob(filepath.Join(dir, "*.go"))
	for _, m := range matches {
		fmt.Println("  match:", filepath.Base(m))
	}
}
