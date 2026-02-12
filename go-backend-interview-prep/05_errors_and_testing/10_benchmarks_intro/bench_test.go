package main

import "testing"

// Package-level var to prevent compiler from optimizing away results.
var result int
var resultStr string

func BenchmarkFib20(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Fib(20)
	}
	result = r
}

func BenchmarkFibIter20(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = FibIter(20)
	}
	result = r
}

func BenchmarkFib10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Fib(10)
	}
	result = r
}

func BenchmarkFibIter10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = FibIter(10)
	}
	result = r
}

func BenchmarkStringConcat100(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StringConcat(100)
	}
	resultStr = r
}
