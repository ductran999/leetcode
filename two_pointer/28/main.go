package main

import "fmt"

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	slow := 0
	for fast := 0; fast < len(haystack); fast++ {
		slow = fast
		if haystack[fast] != needle[0] {
			continue
		}

		// the rest of string not enough char
		if len(haystack)-slow < len(needle) {
			break
		}

		// check substring
		sub := haystack[slow : slow+len(needle)]
		if sub == needle {
			return fast
		}
	}

	return -1
}

func main() {
	h := "abc"
	n := "c"
	fmt.Println(strStr(h, n))
}
