package main

import (
	"fmt"
)

func sortColors(nums []int) {
	l, m, r := 0, 0, len(nums)-1

	for m <= r {
		switch nums[m] {
		case 0:
			nums[m], nums[l] = nums[l], nums[m]
			l++
			m++
		case 1:
			m++
		default:
			nums[m], nums[r] = nums[r], nums[m]
			r--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
