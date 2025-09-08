package main

import (
	"fmt"
	"math/big"
)

func multiply(num1 string, num2 string) string {
	// Parse strings into big.Int
	a := new(big.Int)
	b := new(big.Int)

	a.SetString(num1, 10) // base 10
	b.SetString(num2, 10)

	// Multiply
	result := new(big.Int).Mul(a, b)

	return result.String()
}

func main() {
	fmt.Println(multiply("2", "3"))
}
