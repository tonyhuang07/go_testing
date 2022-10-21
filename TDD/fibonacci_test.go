package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	result := []struct {
		input    int
		expected int
	}{{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34}, {10, 55}}

	var fibo int
	for i, e := range result {
		fibo = Fibonacci2(e.input)
		assert.Equal(t, e.expected, fibo, "Fibonacci(%d) = %d; want %d", i, fibo, e.expected)
	}

}

func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci2(10)
	}
}
