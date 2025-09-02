package main

import "fmt"

type NumArray struct {
	prefixSum []int
	arr       []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{
		prefixSum: []int{nums[0]},
		arr:       nums,
	}

	for i := 1; i < len(nums); i++ {
		sum := nums[i] + n.prefixSum[i-1]
		n.prefixSum = append(n.prefixSum, sum)
	}
	fmt.Println(n.prefixSum)
	return n
}

// (2, 5) = (0, 5) - (0,2) + (2) (2 is overlap)
func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.prefixSum[right]
	}
	return n.prefixSum[right] - n.prefixSum[left] + n.arr[left]
}

func main() {
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(n.SumRange(2, 5))
}
