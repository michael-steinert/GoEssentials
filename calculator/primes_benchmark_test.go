package calculator_test

import (
	"testing"

	"github.com/michael-steinert/GoEssentials/calculator"
)

func BenchmarkIsPrimeWith10000(b *testing.B) {
	benchmarkIsPrime(b, 10_000)
}

func BenchmarkIsPrimeWith100000(b *testing.B) {
	benchmarkIsPrime(b, 100_000)
}

func BenchmarkIsPrimeWith1000000(b *testing.B) {
	benchmarkIsPrime(b, 1_000_000)
}

func benchmarkIsPrime(b *testing.B, value int) {
	// b.ResetTimer()
	// Iteration of Benchmark
	for i := 0; i < b.N; i++ {
		calculator.IsPrime(value)
	}
}
