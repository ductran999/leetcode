package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	res := 0
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for _, val := range freq {
		res += val * (val - 1) / 2
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
}
