package main

import "fmt"

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	arrS := []byte(s)

	l, r := 0, len(s)-1

	for l <= r {
		if isVowel(arrS[l]) && isVowel(arrS[r]) {
			arrS[l], arrS[r] = arrS[r], arrS[l]
			l++
			r--
			continue
		}

		if !isVowel(arrS[l]) {
			l++
		}
		if !isVowel(arrS[r]) {
			r--
		}
	}

	return string(arrS)
}

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}
