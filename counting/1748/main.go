package main

import "fmt"

func sumOfUnique(nums []int) int {
	// Count occurrences of each number
	freq := make(map[int]int, len(nums))
	for _, val := range nums {
		freq[val]++
	}

	// Sum numbers that appear exactly once
	ans := 0
	for num, count := range freq {
		if count == 1 {
			ans += num
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 3, 2}
	fmt.Println(sumOfUnique(nums))
}
