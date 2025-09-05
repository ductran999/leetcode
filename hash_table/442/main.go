package main

import "fmt"

func findDuplicates(nums []int) []int {
	mapNums := make(map[int]int)
	res := []int{}

	for _, val := range nums {
		mapNums[val]++
	}

	for k, val := range mapNums {
		if val > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
