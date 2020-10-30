package arrays

import (
	"reflect"
	"testing"
)

func TestFindThreeLargestNumbers(t *testing.T) {
	type args struct {
		A []int
	}
	testCases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				A: []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7},
			},
			want: []int{18, 141, 541},
		},
		{
			name: "test 2",
			args: args{
				A: []int{55, 7, 8},
			},
			want: []int{7, 8, 55},
		},
		{
			name: "test 3",
			args: args{
				A: []int{55, 43, 11, 3, -3, 10},
			},
			want: []int{11, 43, 55},
		},
		{
			name: "test 4",
			args: args{
				A: []int{7, 8, 3, 11, 43, 55},
			},
			want: []int{11, 43, 55},
		},
		{
			name: "test 5",
			args: args{
				A: []int{55, 7, 8, 3, 43, 11},
			},
			want: []int{11, 43, 55},
		},
		{
			name: "test 6",
			args: args{
				A: []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			},
			want: []int{7, 7, 7},
		},
		{
			name: "test 7",
			args: args{
				A: []int{7, 7, 7, 7, 7, 7, 8, 7, 7, 7, 7},
			},
			want: []int{7, 7, 8},
		},
		{
			name: "test 8",
			args: args{
				A: []int{-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7},
			},
			want: []int{-2, -1, 7},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindThreeLargestNumbers(tc.args.A); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FindThreeLargestNumbers() = %v, want %v", got, tc.want)
			}
		})
	}
}
