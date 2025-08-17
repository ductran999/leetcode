package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		// seen and distance from current to the previous idex <= k
		if idx, ok := seen[num]; ok && i-idx <= k {
			return true
		}
		seen[num] = i
	}

	return false
}

func main() {
	nums := []int{99, 99}
	k := 2

	fmt.Println(containsNearbyDuplicate(nums, k))
}
