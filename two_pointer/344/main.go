package main

import "fmt"

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1

	for l < r {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp
		l++
		r--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Printf("%+v", s)
}
