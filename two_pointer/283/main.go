package main

import "fmt"

func moveZeroes(nums []int) {
	lastNoneZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNoneZero] = nums[lastNoneZero], nums[i]
			lastNoneZero++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
