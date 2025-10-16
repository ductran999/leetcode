package main

import "fmt"

func findSmallestInteger(nums []int, value int) int {
	freq := make([]int, value)
	for _, a := range nums {
		r := a % value
		if r < 0 {
			r += value
		}
		freq[r]++
	}
	x := 0
	for {
		r := x % value
		if freq[r] == 0 {
			return x
		}
		freq[r]--
		x++
	}
}

func main() {
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 5
	fmt.Println(findSmallestInteger(nums, value))
}
