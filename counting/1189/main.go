package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	var freq [26]int
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b', 'a', 'l', 'o', 'n':
			freq[text[i]-'a']++
		}
	}

	counts := []int{
		freq[0],
		freq['b'-'a'],
		freq['n'-'a'],
		freq['o'-'a'] / 2,
		freq['l'-'a'] / 2,
	}

	min := counts[0]
	for _, v := range counts {
		if v < min {
			min = v
		}
	}

	return min
}

func main() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}
