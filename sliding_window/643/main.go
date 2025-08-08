package main

import "fmt"

// fixed window pattern
func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0

	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	// target find the window have total elements is max
	sum := maxSum
	for r := k; r < len(nums); r++ {
		// shift window
		sum += nums[r]
		sum -= nums[r-k]

		// update new max
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Printf("%.2f", findMaxAverage(arr, k))
}
