package main

import "fmt"

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	n := 5
	fmt.Println(canWinNim(n))
}
