package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	arr   []int
	total int
}

func NewStack() *Stack {
	return &Stack{
		arr:   make([]int, 0, 1000),
		total: 0,
	}
}

func (s *Stack) Pop() {
	if len(s.arr) > 0 {
		s.total -= s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
	}
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
	s.total += val
}

func (s *Stack) Peak() int {
	if len(s.arr) > 0 {
		return s.arr[len(s.arr)-1]
	}
	return 0
}

func (s *Stack) Sum() {
	n := len(s.arr)
	total := s.arr[n-1] + s.arr[n-2]
	s.Push(total)
}

func calPoints(operations []string) int {
	s := NewStack()

	for _, op := range operations {
		switch op {
		case "C":
			s.Pop()
		case "D":
			s.Push(s.Peak() * 2)
		case "+":
			s.Sum()
		default:
			d, _ := strconv.Atoi(op)
			s.Push(d)
		}
	}

	return s.total
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
}
