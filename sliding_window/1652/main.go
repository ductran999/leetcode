package main

import "fmt"

func decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, len(code))

	if k == 0 {
		return res
	}

	window := 0
	if k > 0 {
		// init window
		for i := 1; i <= k; i++ {
			window += code[i]
		}
		res[0] = window

		for i := 1; i < len(code); i++ {
			window -= code[i]       // sub value current position
			window += code[(i+k)%n] // add value of end of window
			res[i] = window         // assign window value to current position
		}

		return res
	}

	// init window when k < 0
	for i := n - 1; i > n+k-1; i-- {
		window += code[i]
	}
	res[0] = window

	for i := 1; i < n; i++ {
		window += code[i-1]         // add value prev position
		window -= code[(i+k+n-1)%n] // sub value end of window
		res[i] = window             // assign window value to current position
	}

	return res
}

func main() {
	code, k := []int{2, 4, 9, 3}, -2
	fmt.Println(decrypt(code, k))
}
