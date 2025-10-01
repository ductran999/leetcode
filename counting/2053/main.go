package main

import "fmt"

func kthDistinct(arr []string, k int) string {
	freq := make(map[string]int, len(arr))
	for _, val := range arr {
		freq[val]++
	}

	for _, s := range arr {
		if freq[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}
	return ""
}

func main() {
	arr := []string{"a", "b", "a"}
	k := 3
	fmt.Println(kthDistinct(arr, k))
}
