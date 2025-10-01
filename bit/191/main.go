package main

import (
	"fmt"
)

func hammingWeight(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

func main() {
	fmt.Println(hammingWeight(11))
}
