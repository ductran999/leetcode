package main

import (
	"fmt"
	"slices"
)

func divisorSubstrings(num int, k int) int {
	digits := make([]int, 0)

	x := num
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	slices.Reverse(digits)
	cnt := 0
	for i := 0; i <= len(digits)-k; i++ {
		sub := digits[i : i+k]
		beautyNum := 0
		for _, val := range sub {
			beautyNum = beautyNum*10 + val
		}

		if beautyNum != 0 && num%beautyNum == 0 {
			cnt++
		}
	}

	return cnt
}

func main() {
	num := 430043
	k := 2
	fmt.Println(divisorSubstrings(num, k))
}
