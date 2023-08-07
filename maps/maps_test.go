package maps

import (
	"testing"
)

func BenchmarkMapInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]int)
	}
}

func BenchmarkMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[int]int{1: 1}
		_ = m[1]
	}
}

func BenchmarkMapAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		m[i] = i
	}
}

func BenchmarkMapUpdate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[int]int{1: 1}
		m[1] = 2
	}
}

func BenchmarkMapDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[int]int{1: 1}
		delete(m, 1)
	}
}

func makeLargeMap() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i+1000] = i
	}
	return m
}

func BenchmarkLargeMapGet(b *testing.B) {
	m := makeLargeMap()
	for i := 0; i < b.N; i++ {
		_ = m[1]
	}
}

func BenchmarkLargeMapAdd(b *testing.B) {
	m := makeLargeMap()
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkLargeMapUpdate(b *testing.B) {
	m := makeLargeMap()
	for i := 0; i < b.N; i++ {
		m[1] = 2
	}
}

func BenchmarkLargeMapDelete(b *testing.B) {
	m := makeLargeMap()
	for i := 0; i < b.N; i++ {
		delete(m, 1000)
	}
}
