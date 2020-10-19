package graph

import "testing"

func TestTreeLevelSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    tree
		expected string
	}{
		{
			name: "testcase 1",
			input: tree{
				value: 1,
				left: &tree{
					value: 2,
					left:  &tree{value: 4},
					right: &tree{value: 5},
				},
				right: &tree{
					value: 3,
					right: &tree{
						value: 8,
						left:  &tree{value: 6},
						right: &tree{value: 7},
					},
				},
			},
			expected: "1,5,17,13",
		},
	}
	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			if actual := TreeLevelSum(v.input); actual != v.expected {
				t.Errorf("actual: %v, expected: %v", actual, v.expected)
			}
		})
	}
}
