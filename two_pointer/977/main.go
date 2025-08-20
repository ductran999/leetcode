package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] * nums[i]
	}
	sort.Ints(res)

	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
