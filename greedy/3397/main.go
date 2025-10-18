package main

import (
	"fmt"
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	count := 0
	prev := math.MinInt32 >> 1
	for _, a := range nums {
		low := a - k
		high := a + k
		x := prev + 1
		if x < low {
			x = low
		}
		if x <= high {
			count++
			prev = x
		}
	}
	return count
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 4}
	k := 2
	fmt.Println(maxDistinctElements(nums, k))
}
