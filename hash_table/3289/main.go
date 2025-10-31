package main

import "fmt"

var mapNum [101]int

func getSneakyNumbers(nums []int) []int {
	ans := make([]int, 0, len(nums))

	for _, val := range nums {
		mapNum[val]++
		if mapNum[val] == 2 {
			ans = append(ans, val)
		}
	}
	return ans
}

func main() {
	nums := []int{0, 1, 1, 0}
	fmt.Println(getSneakyNumbers(nums))
}
