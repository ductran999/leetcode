package main

import "fmt"

func maximumLengthSubstring(s string) int {
	freq := make(map[byte]int)

	max := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if len(s)-l <= max {
			break
		}
		freq[s[r]]++
		for freq[s[r]] > 2 {
			freq[s[l]]--
			l++
		}

		if max < r-l+1 {
			max = r - l + 1
		}
	}

	return max
}

func main() {
	s := "eebadadbfa"
	fmt.Println(maximumLengthSubstring(s))
}
