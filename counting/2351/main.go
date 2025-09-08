package main

import "fmt"

func repeatedCharacter(s string) byte {
	chars := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
		if chars[s[i]] > 1 {
			return s[i]
		}
	}
	return 0
}

func main() {
	s := "abccbaacz"
	fmt.Println(repeatedCharacter(s))
}
