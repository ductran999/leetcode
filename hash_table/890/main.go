package main

import "fmt"

func isSamePattern(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq1 := make(map[byte]byte)
	freq2 := make(map[byte]byte)

	for i := 0; i < len(s1); i++ {
		ch1, ch2 := s1[i], s2[i]

		// Check char of string 1 map to string 2
		if val, ok := freq1[ch2]; ok {
			if val != ch1 {
				return false
			}
		} else {
			freq1[ch2] = ch1
		}

		// Check char of string 2 map to string 1
		if val, ok := freq2[ch1]; ok {
			if val != ch2 {
				return false
			}
		} else {
			freq2[ch1] = ch2
		}
	}

	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	ans := []string{}
	for _, word := range words {
		if isSamePattern(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
