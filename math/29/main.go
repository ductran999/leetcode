package main

import (
	"fmt"
)

const (
	MinInt32 = -1 << 31  // -2147483648
	MaxInt32 = 1<<31 - 1 //  2147483647
)

func divide(dividend int, divisor int) int {
	ans := dividend / divisor

	if ans > MaxInt32 {
		return MaxInt32
	}
	if ans < MinInt32 {
		return MinInt32
	}
	return ans
}

func main() {
	divided := -2147483648
	divisor := 1
	fmt.Println(divide(divided, divisor))
}
