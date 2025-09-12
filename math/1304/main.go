package main

import "fmt"

func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}

	init := 1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans[i] = init
		} else {
			ans[i] = -init
			init++
		}
	}

	if n%2 != 0 {
		ans[n-1] = 0
	}

	return ans
}

func main() {
	n := 5
	fmt.Println(sumZero(n))
}
