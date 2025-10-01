package main

import "fmt"

type Stack struct {
	arr []byte
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]byte, 0, 100),
	}
}

func (s *Stack) Push(ch byte) {
	s.arr = append(s.arr, ch)
}

func (s *Stack) Pop() {
	if len(s.arr) > 0 {
		s.arr = s.arr[:len(s.arr)-1]
	}
}

func (s *Stack) Ans() string {
	return string(s.arr)
}

func clearDigits(s string) string {
	stk := NewStack()

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			stk.Pop()
			continue
		}
		stk.Push(s[i])
	}

	return stk.Ans()
}

func main() {
	fmt.Println(clearDigits("cb34"))
}
