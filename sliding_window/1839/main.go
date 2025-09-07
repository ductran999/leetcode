package main

import (
	"fmt"
)

// Check char must  place in right order a e i o u
func isRightVowelOrder(curr, prev byte) bool {
	return (curr == 'e' && prev == 'a') ||
		(curr == 'i' && prev == 'e') ||
		(curr == 'o' && prev == 'i') ||
		(curr == 'u' && prev == 'o')
}

func longestBeautifulSubstring(word string) int {
	res := 0
	if len(word) < 5 {
		return res
	}

	start := 0
	for i := 0; i < len(word); i++ {
		// find start of string must be 'a'
		if word[i] == 'a' {
			start = i
			end := i + 1
			// extend window
			for r := i + 1; r < len(word); r++ {
				end = r
				if word[r] != word[r-1] && !isRightVowelOrder(word[r], word[r-1]) {
					break
				}
				if end == len(word)-1 {
					end++
				}
			}

			// Check string is beautiful: at least 5 chars and end by 'u'.
			// Cause end with u must contain in prev 'a' 'e' 'i' 'o'
			subLen := end - start
			if subLen >= 5 && word[end-1] == 'u' && subLen > res {
				res = subLen
			}

			// Shift window to new pivot
			i = end - 1
		}
	}

	return res
}

func main() {
	w := "aeeeiiiioooauuuaeiou"
	fmt.Println(longestBeautifulSubstring(w))
}
