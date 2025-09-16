package main

import "fmt"

func divideArray(nums []int) bool {
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}

	cntPairs := 0
	for _, f := range freq {
		cntPairs += f / 2
	}

	return cntPairs == len(nums)/2
}

func main() {
	nums := []int{3, 2, 3, 2, 2, 2}
	fmt.Println(divideArray(nums))
}
