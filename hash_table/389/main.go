package main

import "fmt"

func findTheDifference(s string, t string) byte {
	freq1 := make(map[byte]int, len(s))
	freq2 := make(map[byte]int, len(t))

	var res byte

	for i := 0; i < len(s); i++ {
		freq1[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		freq2[t[i]]++
	}

	for k, val := range freq2 {
		if v := freq1[k]; v != val {
			return k
		}
	}

	return res
}

func main() {
	s := "a"
	t := "aa"
	fmt.Println(findTheDifference(s, t))
}
