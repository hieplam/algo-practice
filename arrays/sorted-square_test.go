package arrays

import (
	"reflect"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {

	testCases := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test 1",
			args: []int{1, 2, 3, 5, 6, 8, 9},
			want: []int{1, 4, 9, 25, 36, 64, 81},
		},
		{
			name: "test 2",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "test 3",
			args: []int{1, 2},
			want: []int{1, 4},
		},
		{
			name: "test 4",
			args: []int{1, 2, 3, 4, 5},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "test 5",
			args: []int{0},
			want: []int{0},
		},
		{
			name: "test 6",
			args: []int{-100},
			want: []int{10000},
		},
		{
			name: "test 7",
			args: []int{-2, -1},
			want: []int{1, 4},
		},
		{
			name: "test 8",
			args: []int{-5, -4, -3, -2, -1},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "test 9",
			args: []int{-10, -5, 0, 5, 10},
			want: []int{0, 25, 25, 100, 100},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SortedSquaredArray(tc.args); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SortedSquaredArray() = %v, want %v", got, tc.want)
			}
		})
	}
}
