package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	char := []byte{}
	for i := 0; i < minLength; i++ {
		c := strs[0][i]
		same := true

		for j := 1; j < len(strs); j++ {
			sub := strs[j][i]
			if c != sub {
				same = false
				break
			}
		}
		if same {
			char = append(char, c)
		}
	}

	return string(char)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
