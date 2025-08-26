package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mapIdx := make(map[int]int)

	for i, val := range nums2 {
		mapIdx[val] = i
	}

	ans := make([]int, 0)
	for _, val := range nums1 {
		idx, _ := mapIdx[val]

		found := false
		for i := idx; i < len(nums2); i++ {
			if nums2[i] > val {
				found = true
				ans = append(ans, nums2[i])
				break
			}
		}

		if !found {
			ans = append(ans, -1)
		}
	}

	return ans
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
