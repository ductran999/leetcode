package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)

	n := len(citations)
	if n == 0 {
		return 0
	}

	l, r := 0, n-1

	for l < r {
		// left boundary
		mid := l + (r-l)/2
		if citations[mid] >= n-mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if citations[l] >= n-l {
		return n - l
	}
	return 0
}

func main() {
	ci := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(ci))
}
