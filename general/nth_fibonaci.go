package general

func nthFibonaci(n int) int {
	fibo := []int{0, 1}
	for len(fibo) < n {
		f := fibo[len(fibo)-1] + fibo[len(fibo)-2]
		fibo = append(fibo, f)
	}

	return fibo[n-1]
}
