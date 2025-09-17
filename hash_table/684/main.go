package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	// Track the minimum and maximum lengths of root words in the dictionary
	min, max := len(dictionary[0]), len(dictionary[0])

	// Key optimization 1: Use a map to store root words for O(1) lookup
	dict := make(map[string]bool, len(dictionary))
	for _, root := range dictionary {
		dict[root] = true

		// Update min and max lengths
		n := len(root)
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}

	// Iterate over each word in the sentence
	words := strings.Fields(sentence)
	for idx, word := range words {
		// Key optimization 2: Skip words shorter than the minimum root length
		if len(word) < min {
			continue
		}

		// Key optimization 3: Check prefixes from min to max length
		ceil := len(word)
		if ceil > max {
			ceil = max
		}
		for i := min; i <= ceil; i++ {
			prefix := word[:i]
			if dict[prefix] {
				words[idx] = prefix
				break
			}
		}
	}

	// Return the reconstructed sentence
	return strings.Join(words, " ")
}

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}
