package main

import (
	"fmt"
	"slices"
)

func removeAnagrams(words []string) []string {
	ans := make([]string, 0, len(words))
	prev := ""
	for _, w := range words {
		if w == prev {
			continue
		}

		bytes := []byte(w)
		slices.Sort(bytes)
		sortedWord := string(bytes)

		if sortedWord != prev {
			ans = append(ans, w)
			prev = sortedWord
		}
	}

	return ans
}

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	fmt.Println(removeAnagrams(words))
}
