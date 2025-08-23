package main

import "fmt"

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 4 3 2 1
// i     j
func nextPermutation(nums []int) {
	n := len(nums)

	i := n - 2
	// find pivot: the first element is lower than previous
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		// find right pivot: the first element is greater than i
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}

		// swap i, j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse from i + 1
	reverse(nums, i+1, n-1)
}

func main() {
	nums := []int{4, 3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
