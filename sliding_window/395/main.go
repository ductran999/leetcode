package main

import "fmt"

// window size variable
func longestSubstring(s string, k int) int {
	return helper(s, k)
}

func helper(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	for i, ch := range s {
		if freq[ch] < k {
			left := helper(s[:i], k)
			right := helper(s[i+1:], k)
			if left > right {
				return left
			}
			return right
		}
	}

	return len(s)
}

func main() {
	s := "abacbbdef"
	k := 2
	fmt.Println(longestSubstring(s, k))
}
