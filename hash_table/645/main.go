package main

import "fmt"

func findErrorNums(nums []int) []int {
	res := []int{}
	dups := make(map[int]int)

	for _, val := range nums {
		dups[val]++
	}

	for _, val := range nums {
		if dups[val] > 1 {
			res = append(res, val)
			break
		}
	}

	for i := range nums {
		if dups[i+1] == 0 {
			res = append(res, i+1)
			break
		}
	}

	return res
}

func main() {
	l2 := []int{2, 2}
	fmt.Println(findErrorNums(l2))
}
