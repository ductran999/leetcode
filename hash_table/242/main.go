package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChar := make(map[byte]int, len(s))
	tChar := make(map[byte]int, len(t))

	for i := 0; i < len(s); i++ {
		sChar[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		tChar[t[i]]++
	}

	for c, cnt := range sChar {
		if tCnt, existed := tChar[c]; !existed || tCnt != cnt {
			return false
		}
	}

	return true
}

func main() {
	s, t := "rat", "car"
	fmt.Println(isAnagram(s, t))
}
