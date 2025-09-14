package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	ans := make([][]int, len(matrix))

	for i := 0; i < len(matrix[0]); i++ {
		row := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix); j++ {
			row[j] = matrix[i][j]
		}
		ans[i] = row
	}
	return ans
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(transpose(matrix))
}
