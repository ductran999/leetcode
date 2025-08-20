package main

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func main() {
	nums := []int{1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
