package main

import "fmt"

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		pow := mid * mid

		if pow == num {
			return true
		} else if pow < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	n := 16
	fmt.Println(isPerfectSquare(n))
}
