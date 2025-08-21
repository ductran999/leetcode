package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix[i])-1

		for l <= r {
			mid := l + (r-l)/2

			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	fmt.Println(searchMatrix(matrix, target))
}
