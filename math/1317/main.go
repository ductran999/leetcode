package main

import "fmt"

func isContainZero(n int) bool {
	for n > 0 {
		digit := n % 10
		if digit == 0 {
			return true
		}
		n /= 10
	}
	return false
}

func getNoZeroIntegers(n int) []int {
	res := make([]int, 2)
	for i := 1; i < n; i++ {
		if !isContainZero(i) && !isContainZero(n-i) {
			res[0] = i
			res[1] = n - i
			break
		}
	}

	return res
}

func main() {
	n := 11
	fmt.Println(getNoZeroIntegers(n))
}
