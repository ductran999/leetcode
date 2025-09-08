package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	for _, w := range words1 {
		freq1[w]++
	}
	for _, w := range words2 {
		freq2[w]++
	}

	res := []string{}
	for k, val := range freq1 {
		if val == 1 && freq2[k] == 0 {
			res = append(res, k)
		}
	}
	for k, val := range freq2 {
		if val == 1 && freq1[k] == 0 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2))
}
