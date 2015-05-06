package main

import (
	"math"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(2, 100)
	}
}
