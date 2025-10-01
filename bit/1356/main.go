package main

import (
	"fmt"
	"math/bits"
	"sort"
)

var cacheBit = [10001]int{}

func init() {
	for i := 0; i <= 10000; i++ {
		cacheBit[i] = bits.OnesCount(uint(i))
	}
}

func sortByBits(arr []int) []int {
	sort.Ints(arr) // sort by value first

	buckets := make([][]int, 16) // max 15 bits for numbers ≤ 10000
	for _, val := range arr {
		cnt := cacheBit[val]
		buckets[cnt] = append(buckets[cnt], val)
	}

	ans := make([]int, 0, len(arr))
	for _, b := range buckets {
		ans = append(ans, b...)
	}

	return ans
}

func main() {
	arr := []int{10, 100, 1000, 10000}
	fmt.Println(sortByBits(arr))
}
