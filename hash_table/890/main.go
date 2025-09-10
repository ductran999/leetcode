package main

import "fmt"

// Check if a word follows the same pattern.
// The idea: create a two-way mapping between chars in word and pattern.
// Example: word = "bbc", pattern = "acx"
// b → a
// b → c : inconsistent mapping, so not the same pattern
func isSamePattern(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}

	w2p := make(map[byte]byte) // word char -> pattern char
	p2w := make(map[byte]byte) // pattern char -> word char

	for i := 0; i < len(word); i++ {
		w, p := word[i], pattern[i]

		// Check mapping consistency
		if v, ok := w2p[w]; ok && v != p {
			return false
		}
		if v, ok := p2w[p]; ok && v != w {
			return false
		}

		// Assign mapping if not already set
		w2p[w] = p
		p2w[p] = w
	}

	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	ans := []string{}
	for _, word := range words {
		// append word has same pattern into ans array
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
