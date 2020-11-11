package strings

// LongestPalindromicSubstring accept string s and return longest palindrome substring
func LongestPalindromicSubstring(s string) string {
	r := ""
	for i := range s {
		even := palindromeFrom(s, i, i+1)
		odd := palindromeFrom(s, i-1, i+1)
		if len(r) < len(even) {
			r = even
		}
		if len(r) < len(odd) {
			r = odd
		}
	}
	return r
}

func palindromeFrom(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			return s[l+1 : r]
		}
	}
	return s[l+1 : r]
}
