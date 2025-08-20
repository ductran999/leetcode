package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2] // initialize with first triplet

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			// if this sum is closer, update closest
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closest)) {
				closest = sum
			}

			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				// exact match
				return sum
			}
		}
	}
	return closest
}

func main() {
	nums := []int{0, 0, 0}
	target := 10000
	fmt.Println(threeSumClosest(nums, target))
}
