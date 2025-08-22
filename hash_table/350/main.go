package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}

	map1 := make(map[int]int, len(nums1))
	map2 := make(map[int]int, len(nums2))

	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]]++
	}

	for val, times := range map1 {
		if cnt, existed := map2[val]; existed {
			t := min(cnt, times)
			for t > 0 {
				res = append(res, val)
				t--
			}
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersect(nums1, nums2))
}
