package main

import "fmt"

type Solution struct {
	Dict    map[int][]int
	PickIdx int
}

func Constructor(nums []int) Solution {
	n := len(nums)
	dict := make(map[int][]int, n)
	for i, val := range nums {
		dict[val] = append(dict[val], i)
	}
	return Solution{Dict: dict}
}

func (s *Solution) Pick(target int) int {
	s.PickIdx++
	arr := s.Dict[target]
	return arr[s.PickIdx%len(arr)]
}

func main() {
	nums := []int{1, 2, 3, 3, 3}
	s := Constructor(nums)
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
}
