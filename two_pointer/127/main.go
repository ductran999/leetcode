package main

import (
	"fmt"
	"strings"
)

func isLowerChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1

	for l <= r {
		if !isLowerChar(s[l]) {
			l++
			continue
		}
		if !isLowerChar(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
