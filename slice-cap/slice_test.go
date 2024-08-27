package main

import "testing"

func BenchmarkFromEmptySlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fromEmptySlice()
	}
}

func BenchmarkFromPreAllocatedSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fromPreAllocatedSlice()
	}
}
