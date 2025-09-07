package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// Odd-length palindrome (center is s[i])
		l1, r1 := expand(i, i, s)

		// Even-length palindrome (center is between s[i] and s[i+1])
		l2, r2 := expand(i, i+1, s)

		// Pick the longer palindrome from the two
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// expand expands around the given left (l) and right (r) indexes
func expand(l, r int, s string) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// l and r go one step too far, so return the last valid indices
	return l + 1, r - 1
}

func main() {
	s := "eaabcbaa"
	fmt.Println(longestPalindrome(s))
}
