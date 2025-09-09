package main

import "fmt"

func longestAlternatingSubarray(nums []int, threshold int) int {
	max := 0
	for l := 0; l < len(nums); l++ {
		if nums[l]%2 == 0 && nums[l] <= threshold {
			r := l + 1
			for r < len(nums) {
				if nums[r] > threshold || nums[r]%2 == nums[r-1]%2 {
					break
				}
				r++
			}
			if max < r-l {
				max = r - l
			}
		}
	}

	return max
}

func main() {
	nums := []int{1, 2}
	threshold := 2
	fmt.Println(longestAlternatingSubarray(nums, threshold))
}
