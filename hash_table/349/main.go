package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		freq1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		freq2[nums2[i]]++
	}

	res := []int{}
	for k := range freq1 {
		if v, ok := freq2[k]; ok && v > 0 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
}
