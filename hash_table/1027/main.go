package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}

	dup := make(map[int]struct{})

	for _, val := range freq {
		if _, ok := dup[val]; ok {
			return false
		}
		dup[val] = struct{}{}
	}
	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}
