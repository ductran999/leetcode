package main

import (
	"fmt"
	"slices"
)

func heightChecker(heights []int) int {
	copied := make([]int, 0)
	copied = append(copied, heights...)
	slices.Sort(copied)

	ans := 0
	for i, val := range heights {
		if val != copied[i] {
			ans++
		}
	}

	return ans
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights))
}
