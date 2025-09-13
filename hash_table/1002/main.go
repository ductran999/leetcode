package main

import (
	"fmt"
)

func commonChars(words []string) []string {
	// freq[i][ch] = count of character ch in words[i]
	freq := make([][26]int, len(words))

	// Count frequency
	for i, w := range words {
		for j := 0; j < len(w); j++ {
			freq[i][w[j]-'a']++
		}
	}

	ans := []string{}
	// For each char a-z, find minimum frequency across all words
	for ch := 0; ch < 26; ch++ {
		minCount := freq[0][ch]
		for i := 1; i < len(words); i++ {
			if freq[i][ch] == 0 {
				minCount = 0
				break
			}
			if freq[i][ch] < minCount {
				minCount = freq[i][ch]
			}
		}
		// Add character minCount times
		for i := 0; i < minCount; i++ {
			ans = append(ans, string('a'+byte(ch)))
		}
	}

	return ans
}

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(words))
}
