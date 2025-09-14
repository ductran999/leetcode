package main

import "fmt"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else {
			ans[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return ans
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
