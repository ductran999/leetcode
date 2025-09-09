package main

import "fmt"

type Stack struct {
	arr []byte
}

func (s *Stack) Push(v byte) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() byte {
	if len(s.arr) == 0 {
		return 0
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func isValid(s string) bool {
	stk := Stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk.Push(s[i])
		} else {
			paren := stk.Pop()
			if paren == 0 ||
				(s[i] == ')' && paren != '(') ||
				(s[i] == ']' && paren != '[') ||
				(s[i] == '}' && paren != '{') {
				return false
			}
		}
	}
	return len(stk.arr) == 0
}

func main() {
	s := ")"
	fmt.Println(isValid(s))
}
