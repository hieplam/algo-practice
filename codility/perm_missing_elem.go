package codility

// PermMissingElem solution
// link: https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

func PermMissingElem(A []int) int {
	N := len(A)
	sum := (N + 1) * (N + 2) / 2
	count := 0
	for _, e := range A {
		count += e
	}
	return sum - count
}
