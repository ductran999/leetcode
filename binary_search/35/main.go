package main

import "fmt"

// search exact
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}
