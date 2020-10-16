package codility

import "testing"

func TestConnectedCell(t *testing.T) {

	testCases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "testcase 1",
			input: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 1, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{0, 1, 0, 1, 1},
			},
			expected: 5,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			if actual := ConnectedCell(v.input); actual != v.expected {
				t.Errorf("actual: %v, expected: %v", actual, v.expected)
			}
		})
	}
}
