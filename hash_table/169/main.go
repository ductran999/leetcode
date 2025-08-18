package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	res := 0
	max := -1
	seen := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		seen[val]++

		if cnt, _ := seen[val]; cnt > max {
			max = cnt
			res = val
		}
	}

	return res
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
