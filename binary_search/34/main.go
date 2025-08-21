package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	n := len(nums)
	if n == 0 {
		return res
	}

	// Find left boundary
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return res
	}
	res[0] = left

	// Find right boundary
	left, right = 0, n-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	res[1] = right

	return res
}

func main() {
	nums := []int{2, 2}
	target := 2
	fmt.Println(searchRange(nums, target))
}
