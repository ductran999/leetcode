package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	digitSum := 0
	for n := lowLimit; n > 0; n /= 10 {
		digitSum += n % 10
	}

	// Key optimize 1: using fixed array to reduce time memory allocation
	// the maximum total can be 9 + 9 + 9 + 9 + 9 = 45 so we init an array size 46 (0- -> 45)
	m := [46]int{}
	max := 0
	for i := lowLimit; i <= highLimit; i++ {
		// Update count for digitSum
		m[digitSum]++
		if m[digitSum] > max {
			max = m[digitSum]
		}

		// Key optimized 2: compute next sum from previous sum
		if i < highLimit {
			x := i + 1
			for x%10 == 0 { // handle carry
				digitSum -= 9
				x /= 10
			}
			digitSum++
		}
	}
	return max
}

func main() {
	lowLimit := 1
	highLimit := 10
	fmt.Println(countBalls(lowLimit, highLimit))
}
