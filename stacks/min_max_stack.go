package stacks

import (
	"errors"
	"math"
)

type minMaxStack struct {
	ele []int
}

func new() *minMaxStack {
	ele := minMaxStack{
		ele: []int{},
	}
	return &ele
}

func (s *minMaxStack) push(num int) {
	s.ele = append(s.ele, num)
}

func (s *minMaxStack) pop() (int, error) {
	if len(s.ele) == 0 {
		return 0, errors.New("stack is empty, no value to pop")
	}
	r := s.ele[len(s.ele)-1]
	s.ele = s.ele[:len(s.ele)-1]

	return r, nil
}

func (s *minMaxStack) peek() (int, error) {
	if len(s.ele) == 0 {
		return 0, errors.New("stack is empty, no value to peek")
	}
	return s.ele[len(s.ele)-1], nil
}

func (s *minMaxStack) min() int {
	min := math.MaxInt32
	for _, v := range s.ele {
		if min > v {
			min = v
		}
	}
	return min
}

func (s *minMaxStack) max() int {
	max := math.MinInt32
	for _, v := range s.ele {
		if max < v {
			max = v
		}
	}
	return max
}
