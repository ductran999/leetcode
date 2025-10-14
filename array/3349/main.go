package main

import "fmt"

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if 2*k > n {
		return false
	}

	nextInc := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			nextInc[i] = nextInc[i+1] + 1
		} else {
			nextInc[i] = 0
		}
	}

	need := k - 1
	for i := 0; i+2*k <= n; i++ {
		if nextInc[i] >= need && nextInc[i+k] >= need {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	k := 3
	fmt.Println(hasIncreasingSubarrays(nums, k))
}
