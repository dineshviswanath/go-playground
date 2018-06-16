package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertion_sort_simple(t *testing.T) {
	assert.Equal(t, []int{1, 2}, InsertionSort([]int{2, 1}))
}

func TestInsertion_sort_3_elements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, InsertionSort([]int{1, 3, 2}))
}

func TestInsertion_sort_3_or_more_elements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, InsertionSort([]int{10, 9, 3, 5, 1, 6, 2, 7, 4, 8}))
}

func TestUniqueNumber_first(t *testing.T) {
	assert.Equal(t, float32(0.5), UniqueNumber([]float32{1.0, 0.5, 1.0}))
}

func TestUniqueNumber_last(t *testing.T) {
	assert.Equal(t, float32(5.0), UniqueNumber([]float32{2.0, 2.0, 5.0}))
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 5040, Factorial(7))
}

func TestBinomialCof(t *testing.T) {
	assert.Equal(t, 10, BinomialCof(5, 3))
}

func TestBinomialCofFor3and2(t *testing.T) {
	assert.Equal(t, 10, BinomialCof(3, 2))
}
