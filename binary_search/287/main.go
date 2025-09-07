package main

import (
	"fmt"
)

// This n + 1 size with n value
// so definitely left size or right size will contain the duplicate element
func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		count := 0

		// count how many value <= mid
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		// count > mid mean the duplicate will stay in left
		if count > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}
