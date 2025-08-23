package main

import (
	"fmt"
)

func frequencySort(s string) string {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	bucket := make([][]byte, len(s)+1)
	for k, val := range freq {
		if val > 0 {
			bucket[val] = append(bucket[val], k)
		}
	}

	res := []byte{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for j := 0; j < len(bucket[i]); j++ {
			n := i
			for n > 0 {
				res = append(res, bucket[i][j])
				n--
			}
		}
	}

	return string(res)
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
