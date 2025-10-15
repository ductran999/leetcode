package main

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	increase := func(i int) int {
		inc := 1
		for prev := nums[i-1]; i < n && prev < nums[i]; i++ {
			prev = nums[i]
			inc++
		}
		return inc
	}

	k, previnc := 1, 1
	for i := 1; i < n; i += previnc {
		inc := increase(i)
		k = max(k, min(previnc, inc), inc>>1)
		previnc = inc
	}

	return k
}

func main() {
	nums := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	fmt.Println(maxIncreasingSubarrays(nums))
}
