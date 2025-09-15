package main

import (
	"fmt"
	"strings"
)

func isContainBrokenLetter(word string, brokenMap map[byte]struct{}) bool {
	for i := 0; i < len(word); i++ {
		if _, contains := brokenMap[word[i]]; contains {
			return true
		}
	}
	return false
}

func canBeTypedWords(text string, brokenLetters string) int {
	brokenMap := make(map[byte]struct{})
	for i := 0; i < len(brokenLetters); i++ {
		brokenMap[brokenLetters[i]] = struct{}{}
	}

	// Count number of word can be typed
	cnt := 0
	words := strings.Split(text, " ")
	for _, w := range words {
		if !isContainBrokenLetter(w, brokenMap) {
			cnt++
		}
	}

	return cnt
}

func main() {
	text := "leet code"
	brokenLetters := "e"
	fmt.Println(canBeTypedWords(text, brokenLetters))
}
