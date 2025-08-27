package main

import "fmt"

func findMaxLength(nums []int) int {
	idx := map[int]int{0: -1}

	sum := 0
	maxLen := 0
	for i, val := range nums {
		if val == 0 {
			sum -= 1
		} else {
			sum += val
		}

		if j, ok := idx[sum]; ok {
			// seen this sum before → subarray (j+1 .. i) balanced
			if i-j > maxLen {
				maxLen = i - j
			}
		} else {
			// store first occurrence of this sum
			idx[sum] = i
		}
	}

	return maxLen
}

func main() {
	nums := []int{0, 1, 1, 1, 1, 1, 0, 0, 0}
	fmt.Println(findMaxLength(nums))
}
