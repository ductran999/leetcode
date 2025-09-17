package main

import (
	"fmt"
	"slices"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	ans := make([]int, 0, len(arr1))
	arr1Freq := make(map[int]int, len(arr2))
	arrNonExisted := make([]int, 0, len(arr1)-len(arr2))

	// Step 1: Initialize a map for arr2 to achieve O(1) lookup
	arr2Map := make(map[int]bool, len(arr2))
	for _, val := range arr2 {
		arr2Map[val] = true
	}

	// Step 2: Classify elements of arr1 into two groups
	for _, val := range arr1 {
		if arr2Map[val] { // if number appears in arr2, count its frequency.
			arr1Freq[val]++
		} else { // otherwise, store it in arrNonExisted
			arrNonExisted = append(arrNonExisted, val)
		}
	}

	// Step 3: Sort arrNonExisted in ascending order
	slices.Sort(arrNonExisted)

	// Step 4: Build ans by appending elements from arr1Freq in the order of arr2
	for _, val := range arr2 {
		for range arr1Freq[val] {
			ans = append(ans, val)
		}
	}
	// append arrNonExisted
	return append(ans, arrNonExisted...)
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
