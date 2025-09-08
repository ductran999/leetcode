package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func splitWords(s string) []string {
	s = strings.ToLower(s)
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func mostCommonWord(paragraph string, banned []string) string {
	words := splitWords(paragraph)
	freq := make(map[string]int)

	for _, val := range words {
		freq[val]++
	}

	res := ""
	max := 0

	for k, val := range freq {
		if !slices.Contains(banned, k) && val > max {
			res = k
			max = val
		}
	}

	return res
}

func main() {
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(p, banned))
}
