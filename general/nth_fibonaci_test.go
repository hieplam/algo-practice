package general

import "testing"

func TestNthFibonaci(t *testing.T) {
	tc := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "test case 1",
			want:  5,
			input: 6,
		},
		{
			name:  "test case 2",
			want:  0,
			input: 1,
		},
		{
			name:  "test case 3",
			want:  1,
			input: 2,
		},
		{
			name:  "test case 4",
			want:  1,
			input: 3,
		},
		{
			name:  "test case 5",
			want:  2,
			input: 4,
		},
		{
			name:  "test case 6",
			want:  3,
			input: 5,
		},
		{
			name:  "test case 7",
			want:  8,
			input: 7,
		},
		{
			name:  "test case 8",
			want:  13,
			input: 8,
		},
		{
			name:  "test case 9",
			want:  21,
			input: 9,
		},
		{
			name:  "test case 10",
			want:  34,
			input: 10,
		},
		{
			name:  "test case 11",
			want:  55,
			input: 11,
		},
		{
			name:  "test case 12",
			want:  89,
			input: 12,
		},
		{
			name:  "test case 13",
			want:  144,
			input: 13,
		},
		{
			name:  "test case 14",
			want:  233,
			input: 14,
		},
	}

	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			if actual := nthFibonaci(v.input); actual != v.want {
				t.Errorf("want:%d - actual:%d", v.want, actual)
			}
		})
	}
}
