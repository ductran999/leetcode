package main

import (
	"fmt"
	"sort"
)

// sliding window variable size
func findLHS(nums []int) int {
	// sort the array to ensure is asc.
	sort.Ints(nums)

	res := 0

	l := 0
	for r := 1; r < len(nums); r++ {
		// shrink window if difference > 1
		for l < r && nums[r]-nums[l] > 1 {
			l++
		}

		if res < r-l+1 && nums[r] != nums[l] {
			res = r - l + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// solve by hashMap
func findLHSHashMap(nums []int) int {
	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
	}

	res := 0
	for k, val := range freq {
		if freq[k+1] != 0 {
			l := val + freq[k+1]
			res = max(l, res)
		}
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(findLHSHashMap(nums))
}
