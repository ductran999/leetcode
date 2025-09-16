package main

import "fmt"

// Bit mask for vowel 'a', 'e', 'i', 'o', 'u'
const vowelMask = (1 << 0) | (1 << 4) | (1 << 8) | (1 << 14) | (1 << 20)

func isVowel(ch byte) bool {
	if 'a' <= ch && ch <= 'z' {
		return (1<<(ch-'a'))&vowelMask != 0
	}
	if 'A' <= ch && ch <= 'Z' {
		return (1<<(ch-'A'))&vowelMask != 0
	}
	return false
}

func halvesAreAlike(s string) bool {
	n := len(s)
	half := n / 2
	balance := 0

	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			if i < half {
				balance++
			} else {
				balance--
			}
		}
	}
	return balance == 0
}

func main() {
	s := "UoaUuIEIeaIOuoUUiAaEUIAUAAuEiUIUiUOeUUouIiiaoeiuioiOIosUoEUoIueAoAOUAiiOAUaIOeaoOUuueoOaoXMjkZDIvJlIQzQQUHHeIUZaUgNcflAvNPCTqbrIofxevHndldyTrwBDhLgQssEGmehKiDJLmRZxLzlaoYWQNyqDmU"
	fmt.Println(halvesAreAlike(s))
}
