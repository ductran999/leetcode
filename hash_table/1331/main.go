package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	// Deduplicate numbers in array to reduce sorting time
	seen := make(map[int]struct{}, len(arr))
	keys := make([]int, 0, len(arr))
	for _, val := range arr {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			keys = append(keys, val)
		}
	}

	// Sort unique values
	sort.Ints(keys)

	// Assign rank
	ranks := make(map[int]int, len(keys))
	for i, val := range keys {
		ranks[val] = i + 1
	}

	ans := make([]int, len(arr))
	for i, val := range arr {
		ans[i] = ranks[val]
	}
	return ans
}

func main() {
	arr := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr))
}
