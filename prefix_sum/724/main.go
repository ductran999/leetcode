package main

import "fmt"

// prefix sum O(n) space
func pivotIndex1(nums []int) int {
	preSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		preSum = append(preSum, preSum[i-1]+nums[i])
	}

	if preSum[0] == preSum[len(nums)-1] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if preSum[i-1] == preSum[len(nums)-1]-preSum[i] {
			return i
		}
	}

	return -1
}

// O(1) space optimized
func pivotIndex(nums []int) int {
	// cal total
	total := 0
	for _, v := range nums {
		total += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}

	return -1
}

func main() {
	nums := []int{2, 1, -1}
	fmt.Println(pivotIndex(nums))
}
