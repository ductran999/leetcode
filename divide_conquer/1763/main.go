package main

import (
	"fmt"
)

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	// build set of characters
	set := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		set[s[i]] = true
	}

	for i := 0; i < len(s); i++ {
		hasBadChar := false

		// Check bad character exist
		ch := s[i]
		if ('a' <= ch && ch <= 'z' && !set[ch-32]) || ('A' <= ch && ch <= 'Z' && !set[ch+32]) {
			hasBadChar = true
		}

		// Split string by bad char and check substring
		if hasBadChar {
			l := longestNiceSubstring(string(s[0:i]))
			r := longestNiceSubstring(string(s[i+1:]))
			if len(l) >= len(r) {
				return l
			}
			return r
		}
	}

	return s
}

func main() {
	s := "YazaAay"
	fmt.Println(longestNiceSubstring(s))
}
