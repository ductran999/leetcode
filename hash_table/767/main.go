package main

import (
	"fmt"
	"slices"
)

func reorganizeString(s string) string {
	chars := make(map[rune]int)
	for _, ch := range s {
		chars[ch]++
	}

	track := 0
	var char rune
	charArr := []rune{}
	for ch, freq := range chars {
		if freq > 1 {
			track++
			char = ch
		}
		if track > 1 || freq > 2 {
			return ""
		}
		charArr = append(charArr, ch)
	}
	slices.Sort(charArr)

	charArr = append(charArr, char)
	return string(charArr)
}

func main() {
	s := "aaab"
	fmt.Println(reorganizeString(s))
}
