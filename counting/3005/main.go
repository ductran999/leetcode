package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	// Key optimize: with constraints  1 <= nums.length <= 100
	// We initialize an fixed array with length 101 to hold value [0, 100]
	// to achieve O(1) space and O(1) access.
	// This reduce overhead when allocating with map[int]int
	freq := [101]int{}

	// Step 1: count frequencies
	max := 0
	for _, num := range nums {
		freq[num]++
		if freq[num] > max {
			max = freq[num]
		}
	}

	// Step 2: sum the numbers that have frequency equal to max
	ans := 0
	for _, val := range freq {
		if val == max {
			ans += val
		}
	}
	return ans
}

func main() {
	nums := []int{}
	fmt.Println(maxFrequencyElements(nums))
}
