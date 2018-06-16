package kata

import (
	"sort"
)

// Sorts using insertion method. Not ideal on performance
// Best on simplicity
func InsertionSort(nums []int) []int {
	var temp int
	start := 0
	end := len(nums)
	for i := start + 1; i < end; i++ {
		for j := i; j > start && nums[j-1] > nums[j]; j-- {
			temp = nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = temp
		}
	}
	return nums
}

// To find the unique number in array of floats
// Kata: There is an array with some numbers. All numbers are equal except for one.
// It’s guaranteed that array contains more than 3 numbers.
func UniqueNumber(arr []float32) float32 {

	length := len(arr)
	float64s := make([]float64, length)
	for i, n := range arr {
		float64s[i] = float64(n)
	}
	sort.Float64s(float64s)

	var unique float64
	if float64s[0] == float64s[1] {
		unique = float64s[length-1]
	} else {
		unique = float64s[0]
	}

	return float32(unique)
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func BinomialCof(n int, r int) int {
	return Factorial(n) / (Factorial(n-r) * Factorial(r))
}
