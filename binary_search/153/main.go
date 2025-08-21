package main

import "fmt"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))
}
