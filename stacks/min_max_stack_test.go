package stacks

import (
	"reflect"
	"testing"
)

func minMaxStackTest(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		s := new()
		s.push(5)
		ok := reflect.DeepEqual(s.ele, []int{5})
		if !ok {
			t.Errorf("push fail, want:%d - actual:%d", []int{5}, s.ele)
		}
		s.push(7)
		ok = reflect.DeepEqual(s.ele, []int{5, 7})
		if !ok {
			t.Errorf("push fail, want:%d - actual:%d", []int{5, 7}, s.ele)
		}
		s.push(2)
		ok = reflect.DeepEqual(s.ele, []int{5, 7, 2})
		if !ok {
			t.Errorf("push fail, want:%d - actual:%d", []int{5, 7, 2}, s.ele)
		}
		if p, _ := s.pop(); p != 2 {
			t.Errorf("pop fail, want:%d - actual:%d", 2, p)
		}
		if _, _ = s.pop(); reflect.DeepEqual(s.ele, []int{5}) {
			t.Errorf("pop fail, want:%d - actual:%d", []int{5}, s.ele)
		}
	})
}
