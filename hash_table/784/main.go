package main

import (
	"fmt"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	// Key optimize 1: using fixed array to reduce overhead allocation of map
	needs := [26]int{}
	for i := 0; i < len(licensePlate); i++ {
		ch := licensePlate[i]
		if 'A' <= ch && ch <= 'Z' {
			needs[ch-'A']++
		} else if 'a' <= ch && ch <= 'z' {
			needs[ch-'a']++
		}
	}

	ans := ""
	for _, word := range words {
		// Key optimized 2: only check if potentially shorter
		if ans != "" && len(word) >= len(ans) {
			continue
		}

		chars := [26]int{}
		for i := 0; i < len(word); i++ {
			chars[word[i]-'a']++
		}

		isValid := true
		for ch, need := range needs {
			if chars[ch] < need {
				isValid = false
				break
			}
		}

		if isValid {
			ans = word
		}
	}

	return ans
}

func main() {
	licensePlate := "1s3 456"
	words := []string{"looks", "pest", "stew", "show"}
	fmt.Println(shortestCompletingWord(licensePlate, words))
}
