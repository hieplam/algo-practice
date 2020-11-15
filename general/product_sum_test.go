package general

import "testing"

func TestProductSum(t *testing.T) {
	testcases := []struct {
		name  string
		input []interface{}
		want  int
	}{
		{
			name: "test case 1",
			want: 12,
			input: []interface{}{
				5, 2, []interface{}{7, -1}, 3, []interface{}{6, []interface{}{-13, 8}, 4},
			},
		},
		{
			name:  "test case 2",
			want:  15,
			input: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:  "test case 3",
			want:  18,
			input: []interface{}{1, 2, []interface{}{3}, 4, 5},
		},
		{
			name:  "test case 4",
			want:  27,
			input: []interface{}{[]interface{}{1, 2}, 3, []interface{}{4, 5}},
		},
		{
			name:  "test case 5",
			want:  600,
			input: []interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{5}}}}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if s := ProductSum(tc.input); s != tc.want {
				t.Errorf("want:%d - actual:%d", tc.want, s)
			}
		})
	}
}
