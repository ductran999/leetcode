package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum12 := make(map[int]int)

	for _, s1 := range nums1 {
		for _, s2 := range nums2 {
			sum12[s1+s2]++
		}
	}

	// O(n^2) find possible sum 1 + 2
	res := 0
	for _, s3 := range nums3 {
		for _, s4 := range nums4 {
			need := 0 - s3 - s4
			res += sum12[need]
		}
	}

	return res
}

func main() {
	nums1 := []int{-1, -1}
	nums2 := []int{-1, 1}
	nums3 := []int{-1, 1}
	nums4 := []int{1, -1}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
