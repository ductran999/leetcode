package main

import "fmt"

// fixed window pattern
func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0

	// first window
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	sum := maxSum
	for i := k; i < len(nums); i++ {
		// Shift window to right by subtract left element and add right element
		sum += nums[i]
		sum -= nums[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
