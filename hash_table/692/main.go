package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	// count freq
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	// bucket sort
	buckets := make([][]string, len(words)+1)
	for word, f := range freq {
		buckets[f] = append(buckets[f], word)
	}

	res := []string{}
	for i := len(buckets) - 1; i >= 1; i-- {
		if len(buckets[i]) > 0 {
			sort.Strings(buckets[i])

			for _, val := range buckets[i] {
				res = append(res, val)
				k--
				if k == 0 {
					return res
				}
			}
		}
	}

	return res
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}
