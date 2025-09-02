package main

import (
	"fmt"
	"math"
)

// Hash table approach O(n) -> TLE
func judgeSquareSum1(c int) bool {
	squares := make(map[int]int)

	for i := 0; i <= int(math.Sqrt(float64(c))); i++ {
		if c == i*i || squares[c-i*i] != 0 {
			return true
		}
	}
	return false
}

// Binary search approach
// Core math concept: If c = a^2 + b^2, then both a and b must be <= sqrt(c).
func judgeSquareSum2(c int) bool {

	// O(n.log(n))
	for i := 0; i <= int(math.Sqrt(float64(c))); i++ {

		// O(log(n)) to find the need number
		l, r := 0, int(math.Sqrt(float64(c)))
		for l <= r {
			mid := l + (r-l)/2

			need := c - i*i
			if mid*mid == need {
				return true
			} else if mid*mid > need {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return false
}

// two pointer approach with O(n) complexity and O(1) space
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))

	for l <= r {
		sum := l*l + r*r

		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}

	return false
}

func main() {
	c := 2
	fmt.Println(judgeSquareSum(c))
}
