package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charWord := make(map[byte]string)
	wordChar := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if w, ok := charWord[pattern[i]]; ok && w != words[i] {
			return false
		}
		if c, ok := wordChar[words[i]]; ok && c != pattern[i] {
			return false
		}

		charWord[pattern[i]] = words[i]
		wordChar[words[i]] = pattern[i]
	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog dog dog dog"
	fmt.Println(wordPattern(pattern, s))
}
