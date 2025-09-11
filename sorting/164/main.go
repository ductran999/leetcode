package main

import (
	"fmt"
	"math"
	"sort"
)

func buckSort(nums []int, minVal, maxVal int) []int {
	n := len(nums)
	res := []int{}

	bucketSize := int(math.Ceil(float64(maxVal-minVal) / float64(n-1)))
	bucketCount := (maxVal-minVal)/bucketSize + 1
	buckets := make([][]int, bucketCount)

	// distribute value into bucket
	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		buckets[idx] = append(buckets[idx], num)
	}

	for _, b := range buckets {
		// sort element in bucket O(k*log(k))
		sort.Ints(b)
		res = append(res, b...)
	}

	return res
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	arr := buckSort(nums, minVal, maxVal)

	maxGap := 0
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > maxGap {
			maxGap = arr[i] - arr[i-1]
		}
	}

	return maxGap
}

func main() {
	nums := []int{3, 1}
	fmt.Println(maximumGap(nums))
}
