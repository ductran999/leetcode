package main

import "fmt"

func repeatedNTimes(nums []int) int {
	seen := make(map[int]struct{})

	for _, val := range nums {
		if _, ok := seen[val]; ok {
			return val
		}
		seen[val] = struct{}{}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(nums))
}
