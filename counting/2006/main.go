package main

import "fmt"

// countKDifference returns the number of pairs (i, j)
// such that |nums[i] - nums[j]| == k.
func countKDifference(nums []int, k int) int {
	// Frequency map to count occurrences of each number
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}

	ans := 0
	// For each number, count how many numbers equal (num - k)
	for _, num := range nums {
		if freq[num-k] > 0 {
			ans += freq[num-k]
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))
}
