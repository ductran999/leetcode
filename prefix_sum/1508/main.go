package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	// init sums
	sums := []int{}
	sums = append(sums, nums...)

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			sums = append(sums, nums[i], nums[j])
		}
	}
	fmt.Println(sums)
	sort.Ints(sums)
	ans := 0
	for i := left - 1; left < right; left++ {
		ans += sums[i]
	}

	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	n := 4
	left := 3
	right := 4
	fmt.Println(rangeSum(nums, n, left, right))
}
