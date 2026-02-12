package main

import (
	"fmt"
	"strings"
)

func main() {
	// --- Byte length vs rune count ---
	s := "Hello, world"
	fmt.Println("string:", s)
	fmt.Println("len (bytes):", len(s))
	fmt.Println("rune count:", len([]rune(s)))

	// --- Multi-byte characters ---
	fmt.Println("\n--- Multi-byte ---")
	jp := "Go言語" // "Go" + 2 kanji (3 bytes each)
	fmt.Println("string:", jp)
	fmt.Println("len (bytes):", len(jp))       // 8
	fmt.Println("rune count:", len([]rune(jp))) // 4

	// --- Indexing gives bytes, not runes ---
	fmt.Println("\n--- Byte indexing ---")
	fmt.Printf("jp[0] = %c (byte)\n", jp[0]) // 'G'
	fmt.Printf("jp[2] = %d (byte, not a full rune)\n", jp[2])

	// --- Range iterates by rune ---
	fmt.Println("\n--- Range (by rune) ---")
	for i, r := range jp {
		fmt.Printf("  byte_index=%d rune=%c (U+%04X)\n", i, r, r)
	}
	// Notice: byte_index jumps from 2 to 5 to 8 (kanji are 3 bytes)

	// --- Convert to []rune for character manipulation ---
	fmt.Println("\n--- Rune slice manipulation ---")
	runes := []rune(jp)
	runes[2] = 'X' // replace first kanji
	fmt.Println("modified:", string(runes))

	// --- Convert to []byte ---
	fmt.Println("\n--- Byte slice ---")
	b := []byte("hello")
	b[0] = 'H'
	fmt.Println("modified bytes:", string(b))

	// --- strings.Builder for efficient concatenation ---
	fmt.Println("\n--- strings.Builder ---")
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		fmt.Fprintf(&sb, "item%d ", i)
	}
	fmt.Println("built:", sb.String())

	// --- Useful string functions ---
	fmt.Println("\n--- String functions ---")
	msg := "  Go is Great  "
	fmt.Println("TrimSpace:", "'"+strings.TrimSpace(msg)+"'")
	fmt.Println("Contains:", strings.Contains(msg, "Great"))
	fmt.Println("Split:", strings.Split("a,b,c", ","))
	fmt.Println("Join:", strings.Join([]string{"x", "y", "z"}, "-"))

	// --- Rune literal vs string literal ---
	fmt.Println("\n--- Rune vs string literal ---")
	var r rune = 'A'
	var str string = "A"
	fmt.Printf("rune 'A': type=%T value=%d\n", r, r)
	fmt.Printf("string \"A\": type=%T len=%d\n", str, len(str))
}
