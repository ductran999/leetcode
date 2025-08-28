package main

import "fmt"

func findSubarrays(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}

	sum := make(map[int]int)
	for r := 1; r < len(nums); r++ {
		t := nums[r] + nums[r-1]
		sum[t]++
		if sum[t] >= 2 {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(findSubarrays(nums))
}
