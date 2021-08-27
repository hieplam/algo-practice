package strings

func IsPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	isPalindrome := true
	for l < r {
		if s[l] != s[r] {
			isPalindrome = false
			break
		} else {
			l++
			r--
		}
	}

	return isPalindrome
}
