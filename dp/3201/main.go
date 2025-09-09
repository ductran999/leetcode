package main

import "fmt"

func maximumLength(nums []int) int {
	odds, evens := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			evens++
		} else {
			odds++
		}
	}

	best := max(odds, evens)
	best = max(best, longestAlternating(nums, 0)) // bắt đầu bằng chẵn
	best = max(best, longestAlternating(nums, 1)) // bắt đầu bằng lẻ

	return best
}

func longestAlternating(nums []int, start int) int {
	cnt, expected := 0, start
	for _, v := range nums {
		if v%2 == expected {
			cnt++
			expected ^= 1 // flip 0 <-> 1
		}
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 5, 9, 4, 2}
	fmt.Println(maximumLength(nums))
}
