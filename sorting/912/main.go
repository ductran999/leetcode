package main

import (
	"fmt"
	"math"
	"slices"
)

func sortArray(nums []int) []int {
	minVal, maxVal := nums[0], nums[0]
	for _, val := range nums {
		if minVal > val {
			minVal = val
		}
		if maxVal < val {
			maxVal = val
		}
	}

	// already sort
	if minVal == maxVal {
		return nums
	}

	n := len(nums)
	bucketSize := int(math.Ceil(float64(maxVal-minVal) / float64(n-1)))
	bucketCount := (maxVal-minVal)/bucketSize + 1
	buckets := make([][]int, bucketCount)

	// Distribute value to buckets
	for _, val := range nums {
		idx := (val - minVal) / bucketSize
		buckets[idx] = append(buckets[idx], val)
	}

	ans := make([]int, 0, len(nums))
	for _, b := range buckets {
		slices.Sort(b)
		ans = append(ans, b...)
	}
	return ans
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(nums))
}
