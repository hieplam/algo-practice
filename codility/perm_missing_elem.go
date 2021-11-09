package codility

// PermMissingElem solution
// link: https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
/*
n = 5
1 2 3 4 5 = f
f = 5 + 4 + 3 +2 + 1
f = n + (n-1) + (n-2) + (n-3) + (n-4) + (n-n)
f = 6n + (1 + 2 + 3 + 4 + 5)
f = 6n + f
2f = 6n
2f = (n+1)n
f = (n+1)n / 2
*/
func PermMissingElem(A []int) int {
	N := len(A)
	//since A is missing one number, so actual N wil be N+1
	sum := (N + 1) * (N + 2) / 2 //formulary for 1 + 2 + 3 + ... + N
	count := 0
	for _, e := range A {
		count += e
	}
	return sum - count
}
