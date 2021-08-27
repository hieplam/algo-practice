package strings_test

import (
	"algo_practice/strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testable := []struct {
		Name  string
		Input string
		Want  bool
	}{
		{
			Name:  "tc1",
			Input: "121",
			Want:  true,
		},
	}

	for _, tc := range testable {
		t.Run(tc.Name, func(t *testing.T) {
			if actual := strings.IsPalindrome(tc.Input); actual != tc.Want {
				t.Errorf("want: %v -- actual: %v", tc.Want, actual)
			}
		})
	}
}
