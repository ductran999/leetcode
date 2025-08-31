package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	threshold := len(nums) / 3

	res := make([]int, 0)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	for k, v := range freq {
		if v > threshold {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
