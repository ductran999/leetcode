package main

import "fmt"

func removeElement(nums []int, val int) int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					diff++
					break
				}
			}
		} else {
			diff++
		}
	}

	return diff
}

func main() {
	nums := []int{2}
	val := 3
	fmt.Println(removeElement(nums, val))
}
