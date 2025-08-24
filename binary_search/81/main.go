package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}

		// don't know go left or right cause they equal the safe way is move left++
		// cause target != mid and mid = left so target cannot be left.
		// safe to go
		if nums[l] == nums[mid] {
			l++
		} else if nums[l] < nums[mid] {
			// left <= target < mid
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // right short
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

func main() {
	nums := []int{3, 5, 1}
	target := 3
	fmt.Println(search(nums, target))
}
