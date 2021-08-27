package stacks_test

import (
	"algo_practice/stacks"
	"reflect"
	"testing"
)

func TestMinMaxStack(t *testing.T) {
	//go test -coverprofile cover.out ./... && go tool cover -html=cover.out
	t.Run("simple test", func(t *testing.T) {
		s := stacks.New()
		s.Push(5)
		ok := reflect.DeepEqual(s.Ele, []int{5})
		if !ok {
			t.Errorf("Push fail, want:%d - actual:%d", []int{5}, s.Ele)
		}
		s.Push(7)
		ok = reflect.DeepEqual(s.Ele, []int{5, 7})
		if !ok {
			t.Errorf("Push fail, want:%d - actual:%d", []int{5, 7}, s.Ele)
		}
		s.Push(2)
		ok = reflect.DeepEqual(s.Ele, []int{5, 7, 2})
		if !ok {
			t.Errorf("Push fail, want:%d - actual:%d", []int{5, 7, 2}, s.Ele)
		}
		if p, _ := s.Pop(); p != 2 {
			t.Errorf("Pop fail, want:%d - actual:%d", 2, p)
		}
		if peek, _ := s.Peek(); peek != 7 {
			t.Errorf("Peek fail, want:%d - actual:%d", 7, peek)
		}
		if _, _ = s.Pop(); !reflect.DeepEqual(s.Ele, []int{5}) {
			t.Errorf("Pop fail, want:%d - actual:%d", []int{5}, s.Ele)
		}

	})
}
