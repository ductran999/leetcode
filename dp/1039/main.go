package main

import "fmt"

func minScoreTriangulation(A []int) int {
	n := len(A)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 2; j < n; j++ {
		for i := j - 2; i >= 0; i-- {
			dp[i][j] = int(^uint(0) >> 1) // max int in go
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+A[i]*A[k]*A[j]+dp[k][j])
			}
		}
	}
	return dp[0][n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	values := []int{1, 2, 3}
	fmt.Println(minScoreTriangulation(values))
}
