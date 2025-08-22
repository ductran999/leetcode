package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	res := []int{}

	freq := make(map[int]int, len(nums))
	for _, val := range nums {
		freq[val]++
	}

	// create bucket like hash map using bucket to prevent collision
	bucket := make([][]int, len(nums)+1)
	for k, val := range freq {
		bucket[val] = append(bucket[val], k)
	}

	for i := len(bucket) - 1; i >= 0 && k > len(res); i-- {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
	}

	return res
}

func main() {
	nums := []int{1}
	k := 1

	fmt.Println(topKFrequent(nums, k))
}
