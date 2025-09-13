package main

import "fmt"

func maxFreqSum(s string) int {
	maxCons := 0
	maxVowel := 0

	const vowelMask = (1 << 0) | (1 << 4) | (1 << 8) | (1 << 14) | (1 << 20)

	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		chars[ch]++
		if (1<<ch)&vowelMask != 0 {
			if chars[ch] > maxVowel {
				maxVowel = chars[ch]
			}
		} else {
			if chars[ch] > maxCons {
				maxCons = chars[ch]
			}
		}
	}

	return maxCons + maxVowel
}

func main() {
	s := "successes"
	fmt.Println(maxFreqSum(s))
}
