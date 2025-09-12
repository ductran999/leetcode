package main

import "fmt"

func isInRow(row map[byte]struct{}, word string) bool {
	for i := 0; i < len(word); i++ {
		if _, ok := row[word[i]]; !ok {
			return false
		}
	}
	return true
}

func findWords(words []string) []string {
	ans := []string{}
	row1 := map[byte]struct{}{
		'q': {}, 'w': {}, 'e': {}, 'r': {}, 't': {}, 'y': {}, 'u': {}, 'i': {}, 'o': {}, 'p': {},
		'Q': {}, 'W': {}, 'E': {}, 'R': {}, 'T': {}, 'Y': {}, 'U': {}, 'I': {}, 'O': {}, 'P': {},
	}

	row2 := map[byte]struct{}{
		'a': {}, 's': {}, 'd': {}, 'f': {}, 'g': {}, 'h': {}, 'j': {}, 'k': {}, 'l': {},
		'A': {}, 'S': {}, 'D': {}, 'F': {}, 'G': {}, 'H': {}, 'J': {}, 'K': {}, 'L': {},
	}

	row3 := map[byte]struct{}{
		'z': {}, 'x': {}, 'c': {}, 'v': {}, 'b': {}, 'n': {}, 'm': {},
		'Z': {}, 'X': {}, 'C': {}, 'V': {}, 'B': {}, 'N': {}, 'M': {},
	}

	for _, word := range words {
		if isInRow(row1, word) || isInRow(row2, word) || isInRow(row3, word) {
			ans = append(ans, word)
		}
	}

	return ans
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}
