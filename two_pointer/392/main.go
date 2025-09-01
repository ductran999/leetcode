package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return false
	}

	slow := 0
	for fast := 0; fast < len(t); fast++ {
		if t[fast] == s[slow] {
			slow++
		}
		if slow >= len(s) {
			break
		}
	}
	return slow-len(s) == 0
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
