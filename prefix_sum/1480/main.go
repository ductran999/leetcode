package main

import "fmt"

func runningSum(nums []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
