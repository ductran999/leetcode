package main

import "fmt"

func sumZero(n int) []int {
	ans := make([]int, n)

	// Fill the array with pairs (+x, -x)
	val := 1
	for i := 0; i+1 < n; i += 2 {
		ans[i] = val    // positive
		ans[i+1] = -val // negative
		val++
	}

	// If n is odd, set the last element to 0 to balance the sum
	if n%2 != 0 {
		ans[n-1] = 0
	}

	return ans
}

func main() {
	n := 5
	fmt.Println(sumZero(n))
}
