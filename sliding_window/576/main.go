package main

import "fmt"

func isMatch(a map[byte]int, b map[byte]int) bool {
	for k, v := range a {
		if currVal, ok := b[k]; !ok || currVal != v {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	k := len(s1)
	required := make(map[byte]int)
	has := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		required[s1[i]]++
		has[s2[i]]++
	}

	if isMatch(required, has) {
		return true
	}

	for right := k; right < len(s2); right++ {
		has[s2[right]]++
		has[s2[right-k]]--

		if isMatch(required, has) {
			return true
		}
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidboaoo"

	fmt.Println(checkInclusion(s1, s2))
}
