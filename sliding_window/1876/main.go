package main

import "fmt"

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	cnt := 0
	for i := 2; i < len(s); i++ {
		if s[i]-s[i-1] != 0 && s[i-1]-s[i-2] != 0 && s[i] != s[i-2] {
			cnt++
		}
	}
	return cnt
}

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings(s))
}
