package main

import (
	"fmt"
)

// longest substrings without not contain duplicate chars
// variable window size narrow left when got duplicates
// Approach: tracking number of chars
func lengthOfLongestSubstring1(s string) int {
	res := 0
	left := 0

	chMap := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		chMap[s[r]]++

		// duplicate then narrow until it not duplicate
		for chMap[s[r]] > 1 {
			chMap[s[left]]--
			left++
		}

		// update new length
		if r-left+1 > res {
			res = r - left + 1
		}
	}

	return res
}

// Approach keep track by last index
func lengthOfLongestSubstring2(s string) int {
	res := 0
	left := 0

	seen := make(map[byte]int, len(s))
	for right := 0; right < len(s); right++ {
		// check char existed and update new left
		if idx, existed := seen[s[right]]; existed && idx >= left {
			left = idx + 1
		}

		// update new index for this character
		seen[s[right]] = right

		// update the new longest length
		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring1(s))
	fmt.Println(lengthOfLongestSubstring2(s))
}
