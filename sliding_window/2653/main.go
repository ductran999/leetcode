package main

import (
	"fmt"
	"sort"
)

func getSmallTh(mapNums map[int]int, x int) int {
	// Not contain negative
	if len(mapNums) == 0 {
		return 0
	}
	keys := make([]int, 0)

	for k, val := range mapNums {
		if val > 0 {
			keys = append(keys, k)
		}
	}

	sort.Ints(keys)
	cnt := 0
	for _, num := range keys {
		cnt += mapNums[num]
		if cnt >= x {
			return num
		}
	}

	return 0
}

// fix size window pattern
func getSubarrayBeauty(nums []int, k int, x int) []int {
	res := []int{}
	freq := make(map[int]int)

	// init window
	for i := range k {
		if nums[i] < 0 {
			freq[nums[i]]++
		}
	}
	res = append(res, getSmallTh(freq, x))

	for r := k; r < len(nums); r++ {
		if nums[r] < 0 {
			freq[nums[r]]++
		}
		if nums[r-k] < 0 {
			freq[nums[r-k]]--
		}
		res = append(res, getSmallTh(freq, x))
	}
	return res
}

func main() {
	nums := []int{1, -1, -3, -2, 3}
	k := 3
	x := 2
	fmt.Println(getSubarrayBeauty(nums, k, x))
}
