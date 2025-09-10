package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}

	freq1 := make(map[byte]int)
	freq2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq1[s[i]]++
	}
	for i := 0; i < len(goal); i++ {
		freq2[goal[i]]++
	}

	if len(freq1) != len(freq2) {
		return false
	}
	if len(freq1) == len(freq2) && len(freq1) == 1 {
		return true
	}
	for k, v := range freq1 {
		if freq2[byte(k)] != v {
			return false
		}
	}

	overone := false
	for _, val := range freq1 {
		if val > 1 {
			overone = true
		}
	}

	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			cnt++
		}
		if cnt > 2 {
			return false
		}
	}

	// string contains all single char
	if !overone && cnt != 2 {
		return false
	}

	return true
}

func main() {
	s := "aaaaaaabc"
	goal := "aaaaaaacb"
	fmt.Println(buddyStrings(s, goal))
}
