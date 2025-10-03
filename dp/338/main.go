package main

import "fmt"

var bits [100001]int

func init() {
	for i := 0; i <= 100000; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
}

func countBits(n int) []int {
	return bits[:n+1]
}

func main() {
	fmt.Println(countBits(5))
}
