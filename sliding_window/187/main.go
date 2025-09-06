package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	ans := []string{}
	if len(s) < 10 {
		return ans
	}

	adnMap := make(map[string]int)
	chars := []byte(s)

	// init window
	window := []byte{}
	for i := 0; i < 10; i++ {
		window = append(window, chars[i])
	}
	adnMap[string(window)]++

	for i := 1; i < len(chars)-9; i++ {
		sub := chars[i : i+10]
		adnMap[string(sub)]++
	}

	for k, v := range adnMap {
		if v > 1 {
			ans = append(ans, k)
		}
	}

	return ans
}

func main() {
	s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
