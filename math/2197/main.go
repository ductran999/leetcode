package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimes(nums []int) []int {
	stack := make([]int, 0, len(nums))

	for _, num := range nums {
		stack = append(stack, num) // push new value

		// Check number in stack ensure their gcd must be 1
		for len(stack) > 1 {
			n := len(stack)
			if gcd(stack[n-1], stack[n-2]) == 1 {
				break
			}

			stack[n-2] = lcm(stack[n-1], stack[n-2])
			stack = stack[:n-1] // pop
		}
	}
	return stack
}

func main() {
	nums := []int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825}
	fmt.Println(replaceNonCoprimes(nums))
}
