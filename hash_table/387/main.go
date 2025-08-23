package main

import "fmt"

func firstUniqChar(s string) int {
	mChars := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mChars[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if cnt := mChars[s[i]]; cnt == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
