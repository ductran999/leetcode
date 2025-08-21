package main

import "fmt"

// peak -> left boundary
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
