package main

import (
	"fmt"
	"strings"
)

func maxRepeating(sequence string, word string) int {
	count := 0
	tmp := word
	for strings.Contains(sequence, tmp) {
		count++
		tmp += word
	}
	return count
}

func main() {
	sequence := "aaaabab"
	word := "aa"
	fmt.Println(maxRepeating(sequence, word))
}
