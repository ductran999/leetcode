package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
