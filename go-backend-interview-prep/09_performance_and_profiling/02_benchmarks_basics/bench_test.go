package main

import (
	"strings"
	"testing"
)

// Prevent dead code elimination: assign result to package-level var.
var result string

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 1000; j++ {
			s += "a"
		}
		result = s
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.Grow(1000)
		for j := 0; j < 1000; j++ {
			sb.WriteByte('a')
		}
		result = sb.String()
	}
}

// BenchmarkJoin shows strings.Join for comparison.
func BenchmarkJoin(b *testing.B) {
	parts := make([]string, 1000)
	for i := range parts {
		parts[i] = "a"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = strings.Join(parts, "")
	}
}
