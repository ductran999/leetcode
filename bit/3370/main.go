package main

import "fmt"

func smallestNumber(n int) int {
	for n&(n+1) != 0 {
		n |= n + 1
	}
	return n
}

func main() {
	fmt.Println(smallestNumber(5))
}
