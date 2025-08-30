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

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(findLHS(nums))
}
