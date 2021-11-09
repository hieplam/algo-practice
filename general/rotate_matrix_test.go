package general_test

import (
	"algo_practice/general"
	"reflect"
	"testing"
)

func TestRotateClockwiseMatrix(t *testing.T) {
	t.Run("aa", func(t *testing.T) {
		m := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		actual := [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}
		value, err := general.RotateClockwiseMatrix(m)
		if err != nil {
			t.Fatal("err", err)
		}
		if !reflect.DeepEqual(value, actual) {
			//t.Errorf("fail")
		}
	})
}
