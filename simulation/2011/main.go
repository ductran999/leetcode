package main

import "fmt"

func finalValueAfterOperations(operations []string) (x int) {
	for _, op := range operations {
		if op == "X++" || op == "++X" {
			x++
		} else {
			x--
		}
	}
	return
}

func main() {
	ops := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(ops))
}
