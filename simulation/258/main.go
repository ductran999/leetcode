package main

import "fmt"

// Any number can be expressed in base 10 as:
// num = d_k * 10^k + d_(k-1) * 10^(k-1) + ... + d_1 * 10 + d_0
// If we mod 9
// 10 ≡ 1 (mod 9)
// 10^2 ≡ 1 (mod 9)
// num % 9 = (d_k + d_(k-1) + ... + d_1 + d_0) % 9
// Example:
// 38     = 3 * 10^1 + 8 * 10^0
//        = 3 * 10 + 8 * 1
// 38 % 9 = (3 + 8) % 9
//        = 11 % 9
//        = 2
// But in case num = 9:
// 9 % 9 = 0  → expected digital root = 9
// Trick:
// If num == 0 → return 0
// Else → return 1 + (num - 1) % 9
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + (num-1)%9
}

func main() {
	n := 38
	fmt.Println(addDigits(n))
}
