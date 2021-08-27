package stacks

import (
	"errors"
	"math"
)

type MinMaxStack struct {
	Ele []int
}

func New() *MinMaxStack {
	ele := MinMaxStack{
		Ele: []int{},
	}
	return &ele
}

func (s *MinMaxStack) Push(num int) {
	s.Ele = append(s.Ele, num)
}

func (s *MinMaxStack) Pop() (int, error) {
	if len(s.Ele) == 0 {
		return 0, errors.New("stack is empty, no value to Pop")
	}
	r := s.Ele[len(s.Ele)-1]
	s.Ele = s.Ele[:len(s.Ele)-1]

	return r, nil
}

func (s *MinMaxStack) Peek() (int, error) {
	if len(s.Ele) == 0 {
		return 0, errors.New("stack is empty, no value to Peek")
	}
	return s.Ele[len(s.Ele)-1], nil
}

func (s *MinMaxStack) Min() int {
	min := math.MaxInt32
	for _, v := range s.Ele {
		if min > v {
			min = v
		}
	}
	return min
}

func (s *MinMaxStack) Max() int {
	max := math.MinInt32
	for _, v := range s.Ele {
		if max < v {
			max = v
		}
	}
	return max
}
