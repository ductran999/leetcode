package main

import (
	"fmt"
	"slices"
)

func addBinary(a, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		res = append(res, byte(sum%2)+'0')
		carry = sum / 2
	}

	slices.Reverse(res)
	return string(res)
}

func main() {
	fmt.Println(addBinary("110111", "101"))
}
