package arrays

func SortedSquaredArray(a []int) []int {
	start := 0
	end := len(a) - 1
	result := make([]int, len(a))
	i := end
	for start <= end {
		if abs(a[start]) > abs(a[end]) {
			result[i] = a[start] * a[start]
			start++
		} else {
			result[i] = a[end] * a[end]
			end--
		}
		i--
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
