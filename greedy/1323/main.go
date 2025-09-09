package main

import "fmt"

func maximum69Number(num int) int {
	arr := make([]int, 0)
	for num > 0 {
		digit := num % 10
		arr = append(arr, digit)
		num /= 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 6 {
			arr[i] = 9
			break
		}
	}

	ans := 0
	for i := len(arr) - 1; i >= 0; i-- {
		ans = ans*10 + arr[i]
	}

	return ans
}

func main() {
	num := 9669
	fmt.Println(maximum69Number(num))
}
