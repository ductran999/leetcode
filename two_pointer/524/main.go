package main

import (
	"fmt"
	"sort"
)

func isSubsequence(sub, s string) bool {
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		if s[fast] == sub[slow] {
			slow++
			if slow == len(sub) {
				return true
			}
		}
	}
	return false
}

func findLongestWord(s string, dictionary []string) string {
	res := ""
	sort.Strings(dictionary)

	// check for each word in dictionary
	for _, word := range dictionary {
		// new word valid and more length than the previous update the res
		if isSubsequence(word, s) && len(word) > len(res) {
			res = word
		}
	}

	return res
}

func main() {
	s := "abce"
	dictionary := []string{"abe", "abc"}
	fmt.Println(findLongestWord(s, dictionary))
}
