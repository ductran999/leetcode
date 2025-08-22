package main

import "fmt"

func missingNumber(nums []int) int {
	cnt := make([]int, len(nums)+1)
	for i := range nums {
		cnt[nums[i]] = 1
	}

	for k, v := range cnt {
		if v == 0 {
			return k
		}
	}

	return -1
}

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
}
