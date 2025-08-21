package main

import "fmt"

// find exact
func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		mid := l + (r-l)/2
		pow := mid * mid

		if pow == x {
			return mid
		} else if pow > x {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return r
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}
