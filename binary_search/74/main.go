package main

import "fmt"

// Exact match pattern
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
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 13

	fmt.Println(searchMatrix(matrix, target))
}
