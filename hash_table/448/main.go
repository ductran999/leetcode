package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	res := []int{}

	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
	}

	for i := 1; i <= len(nums); i++ {
		if freq[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
