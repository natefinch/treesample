package main

import (
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		t := NewStringTree([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
		b.StartTimer()
		t.Insert("pineapple")
	}
}

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		t := NewStringTree([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
		b.StartTimer()
		t.Find("monkey")
	}
}
