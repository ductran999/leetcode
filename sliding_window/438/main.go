package main

import "fmt"

func isEqual(a map[byte]int, b map[byte]int) bool {
	for k, v := range a {
		if currVal, ok := b[k]; !ok || v != currVal {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	res := []int{}

	k := len(p)
	if len(s) < len(p) {
		return res
	}

	required := make(map[byte]int, len(p))
	has := make(map[byte]int, len(p))

	for i := 0; i < k; i++ {
		required[p[i]]++
		has[s[i]]++
	}

	if isEqual(required, has) {
		res = append(res, 0)
	}

	for right := len(p); right < len(s); right++ {
		// narrow when shift
		has[s[right]]++
		has[s[right-k]]--

		if isEqual(required, has) {
			res = append(res, right-k+1)
		}
	}

	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
