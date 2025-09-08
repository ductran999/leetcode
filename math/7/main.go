package main

import "fmt"

func reverse(x int) int {
	l := -((1 << 31) - 1)
	r := ((1 << 31) - 1)
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}

	ans := 0
	for x > 0 {
		digit := x % 10
		x /= 10
		ans = (ans*10 + digit)
	}

	if negative {
		if -ans < l {
			return 0
		}
		return -ans
	}

	if ans >= r {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(reverse(-123))
}
