package util

import "math/rand"

func GenerateCase(n int) []int {
	if n < 0 {
		n = 100
	}
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		j := int(rand.Float64() * 500)
		arr = append(arr, j)
	}
	return arr
}
