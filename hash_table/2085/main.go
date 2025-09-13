package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	w1 := make(map[string]int)
	w2 := make(map[string]int)

	for _, w := range words1 {
		w1[w]++
	}
	for _, w := range words2 {
		if w1[w] > 0 {
			w2[w]++
		}
	}

	cnt := 0
	for word, freq := range w1 {
		if w2[word] == freq && freq == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	words1 := []string{"leetcode", "is", "amazing", "as", "is"}
	words2 := []string{"amazing", "leetcode", "is"}
	fmt.Println(countWords(words1, words2))
}
