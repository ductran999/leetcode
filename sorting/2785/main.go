package main

import (
	"fmt"
)

func sortVowels(s string) string {
	// Sorted order of vowels
	order := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

	// Step 1: Count vowels
	var freq [118]int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isVowel(ch) {
			freq[ch]++
		}
	}

	// Step 2: Replace vowels in sorted order
	b := []byte(s)
	pos := 0 // index in `order`

	for i := 0; i < len(b); i++ {
		if isVowel(b[i]) {
			// Find next smallest vowel with count > 0
			for pos < len(order) && freq[order[pos]] == 0 {
				pos++
			}
			b[i] = order[pos]
			freq[order[pos]]--
		}
	}
	return string(b)
}

func isVowel(c byte) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func main() {
	s := "lEetcOde"
	fmt.Println(sortVowels(s))
}
