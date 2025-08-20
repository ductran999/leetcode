package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		// already sort so if sum < target is mean left number + max < target
		// so left number is too small so move to next left number
		if sum < target {
			l++
		} else {
			r--
		}
	}

	return nil
}

func main() {
	nums := []int{-1, 0}
	target := -1
	fmt.Println(twoSum(nums, target))
}
