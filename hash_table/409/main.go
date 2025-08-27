package main

import "fmt"

func longestPalindrome(s string) int {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	odd := 0
	length := 0
	for _, v := range freq {
		length += (v / 2) * 2
		if v%2 == 1 {
			odd = 1
		}
	}

	return length + odd
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
}
