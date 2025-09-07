package main

import "fmt"

func arrangeCoins(n int) int {
	coin := n
	for i := 1; i <= n; i++ {
		if coin < i {
			return i - 1
		}
		coin -= i
	}

	return n
}

func main() {
	n := 8
	fmt.Println(arrangeCoins(n))
}
