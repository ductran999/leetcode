package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	rChar := make(map[byte]int)
	mChar := make(map[byte]int)

	for i := 0; i < len(ransomNote); i++ {
		rChar[ransomNote[i]]++
	}
	for i := 0; i < len(magazine); i++ {
		mChar[magazine[i]]++
	}

	for k, v := range rChar {
		if has := mChar[k]; has < v {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"

	fmt.Println(canConstruct(ransomNote, magazine))
}
