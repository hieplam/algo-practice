package strings

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test case 1",
			input: "abaxyzzyxf",
			want:  "xyzzyx",
		},
		{
			name:  "test case 2",
			input: "a",
			want:  "a",
		},
		{
			name:  "test case 3",
			input: "it's highnoon",
			want:  "noon",
		},
		{
			name:  "test case 4",
			input: "noon high it is",
			want:  "noon",
		},
		{
			name:  "test case 5",
			input: "abccbait's highnoon",
			want:  "abccba",
		},
		{
			name:  "test case 6",
			input: "abcdefgfedcbazzzzzzzzzzzzzzzzzzzz",
			want:  "zzzzzzzzzzzzzzzzzzzz",
		},
		{
			name:  "test case 7",
			input: "abcdefgfedcba",
			want:  "abcdefgfedcba",
		},
		{
			name:  "test case 8",
			input: "abcdefghfedcbaa",
			want:  "aa",
		},
		{
			name:  "test case 9",
			input: "abcdefggfedcba",
			want:  "abcdefggfedcba",
		},
		{
			name:  "test case 10",
			input: "zzzzzzz2345abbbba5432zzbbababa",
			want:  "zz2345abbbba5432zz",
		},
		{
			name:  "test case 11",
			input: "z234a5abbbba54a32z",
			want:  "5abbbba5",
		},
		{
			name:  "test case 12",
			input: "z234a5abbba54a32z",
			want:  "5abbba5",
		},
		{
			name:  "test case 13",
			input: "ab12365456321bb",
			want:  "b12365456321b",
		},
		{
			name:  "test case 14",
			input: "aca",
			want:  "aca",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := LongestPalindromicSubstring(tc.input); actual != tc.want {
				t.Errorf("expected:%s - actual:%s", tc.want, actual)
			}
		})
	}
}
