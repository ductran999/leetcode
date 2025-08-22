package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// using hashmap to detect duplicate
	numMap := make(map[int]struct{}, len(nums))

	for _, val := range nums {
		if _, ok := numMap[val]; ok {
			return true
		}
		numMap[val] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
