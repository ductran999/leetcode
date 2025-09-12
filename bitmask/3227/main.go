package main

import (
	"fmt"
)

// Alice wins if the string contains at least one vowel
func doesAliceWin(s string) bool {
	// Create a bitmask representing vowels:
	// 'a' → 1 << 0   (bit 0)  'a' - 'a' = 0
	// 'e' → 1 << 4   (bit 4)  'e' - 'a' = 4
	// 'i' → 1 << 8   (bit 8)  'i' - 'a' = 8
	// 'o' → 1 << 14  (bit 14) 'o' - 'a' = 14
	// 'u' → 1 << 20  (bit 20) 'u' - 'a' = 20
	// Combine all with OR to get a single mask with all vowel bits set
	const vowelMask = (1 << 0) | (1 << 4) | (1 << 8) | (1 << 14) | (1 << 20)

	for i := 0; i < len(s); i++ {
		// Compute bit for current character: 1 << (s[i] - 'a')
		// & with vowelMask to check if it's a vowel
		if vowelMask&(1<<(s[i]-'a')) != 0 {
			return true
		}
	}
	return false
}

func main() {
	s := "leetcoder"
	fmt.Println(doesAliceWin(s))
}
