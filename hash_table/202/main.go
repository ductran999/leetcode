package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		n = sum
	}

	return n == 1
}

func main() {
	fmt.Println(isHappy(7))
}
