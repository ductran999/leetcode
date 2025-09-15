package main

import "fmt"

// findFirstPositive returns the index of the first element > 0
func findFirstPositive(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= 0 { // still <= 0 → move right
			l = mid + 1
		} else { // found candidate > 0 → shrink left
			r = mid - 1
		}
	}

	// l is the first index with nums[l] > 0 (or len(nums) if none)
	return l
}

// findNonNegative returns the index of the first element >= 0
func findNonNegative(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < 0 { // still negative → move right
			l = mid + 1
		} else { // found candidate >= 0 → shrink left
			r = mid - 1
		}
	}

	// l is the first index with nums[l] >= 0 (or len(nums) if none)
	return l
}

// maximumCount returns the maximum count of either positive or negative numbers
func maximumCount(nums []int) int {
	n := len(nums)

	// If all numbers are positive or all are negative
	if nums[0] > 0 || nums[n-1] < 0 {
		return n
	}

	// l = count of negatives (first index where nums[i] >= 0)
	// r = index where positives start (first nums[i] > 0)
	l := findNonNegative(nums)
	r := findFirstPositive(nums)

	negatives := l
	positives := n - r

	if negatives > positives {
		return negatives
	}
	return positives
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2}
	fmt.Println(maximumCount(nums))
}
