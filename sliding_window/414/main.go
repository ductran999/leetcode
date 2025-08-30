package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)

	max := nums[len(nums)-1]
	count := 2
	res := nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i-- {
		if count == 0 {
			break
		}

		if nums[i] != res {
			count--
			res = nums[i]
		}
	}

	if count == 0 {
		return res
	}

	return max
}

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(thirdMax(nums))
}
