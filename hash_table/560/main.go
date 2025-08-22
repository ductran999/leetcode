package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := map[int]int{0: 1}
	sum := 0
	res := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		// looking for number of sub array that sum up = sum -k
		res += prefixSum[sum-k]

		// track prefixSum
		prefixSum[sum]++
	}

	return res
}

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}
