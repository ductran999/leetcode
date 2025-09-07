package main

import "fmt"

func findDuplicate(nums []int) int {
	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
		if freq[val] > 1 {
			return val
		}
	}

	return 0
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
