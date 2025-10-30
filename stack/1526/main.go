package main

import "fmt"

func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}
	ans := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			ans += target[i] - target[i-1]
		}
	}
	return ans
}

func main() {
	target := []int{1, 2, 3, 2, 1}
	fmt.Println(minNumberOperations(target))
}
