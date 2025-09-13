package main

import "fmt"

func countCharacters(words []string, chars string) int {
	mapChars := [26]int{}

	for i := 0; i < len(chars); i++ {
		mapChars[chars[i]-'a']++
	}

	ans := 0
	for _, w := range words {
		wChars := [26]int{}

		invalid := false
		for i := 0; i < len(w); i++ {
			ch := w[i] - 'a'
			wChars[ch]++

			if mapChars[ch] == 0 || wChars[ch] > mapChars[ch] {
				invalid = true
				break
			}
		}

		if !invalid {
			ans += len(w)
		}
	}

	return ans
}

func main() {
	chars := "atach"
	words := []string{"cat", "bt", "hat", "tree"}
	fmt.Println(countCharacters(words, chars))
}
