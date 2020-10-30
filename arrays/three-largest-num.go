package arrays

import (
	"math"
)

// FindThreeLargestNumbers Find three largest numbers of the slice
func FindThreeLargestNumbers(a []int) []int {
	result := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, v := range a {
		if result[2] < v {
			insertAndShift(result, 2, v)
		} else if result[1] < v {
			insertAndShift(result, 1, v)
		} else if result[0] < v {
			insertAndShift(result, 0, v)
		}
	}
	return result
}

func insertAndShift(a []int, p int, v int) {
	if p == 0 {
		a[0] = v
	} else {
		for i := 1; i <= p; i++ {
			a[i-1] = a[i]
		}
		a[p] = v
	}
}
