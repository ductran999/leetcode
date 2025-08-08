package main

import "fmt"

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	l := 0
	required := make(map[byte]int)
	windowCounts := make(map[byte]int)
	for i := range len(t) {
		required[t[i]]++
	}

	have, need := 0, len(required) // keep track char in window
	resLen, resStart := len(s)+1, 0

	for r := 0; r < len(s); r++ {
		// extend the window
		ch := s[r]
		windowCounts[ch]++

		if windowCounts[ch] == required[ch] {
			have++
		}

		for have == need {
			// Cập nhật kết quả nếu window nhỏ hơn
			if r-l+1 < resLen {
				resLen = r - l + 1
				resStart = l
			}

			// Thu hẹp window từ trái
			windowCounts[s[l]]--
			if windowCounts[s[l]] < required[s[l]] {
				have--
			}
			l++
		}
	}

	if resLen == len(s)+1 {
		return ""
	}

	return s[resStart : resStart+resLen]
}

func main() {
	s := "ADOBBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}
