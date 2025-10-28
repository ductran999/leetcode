package main

import "fmt"

func simulate(nums []int, start int, dir int) bool {
	n := len(nums)
	a := make([]int, n)
	copy(a, nums)
	curr := start
	for curr >= 0 && curr < n {
		if a[curr] == 0 {
			curr += dir
		} else {
			a[curr]--
			dir = -dir
			curr += dir
		}
	}
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func countValidSelections(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			continue
		}
		if simulate(nums, i, -1) {
			ans++
		}
		if simulate(nums, i, 1) {
			ans++
		}
	}
	return ans
}

func main() {
	nums := []int{1, 0, 2, 0, 3}
	fmt.Println(countValidSelections(nums))
}
