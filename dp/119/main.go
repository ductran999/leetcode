package main

import "fmt"

var triangle = make([][]int, 34)

func init() {
	triangle[0] = []int{1}

	for i := 1; i < 34; i++ {
		row := make([]int, i+1)

		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle[i] = row
	}
}

func getRow(rowIndex int) []int {
	return triangle[rowIndex]
}

func main() {
	fmt.Println(getRow(33))
}
