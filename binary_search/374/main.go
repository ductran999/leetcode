package main

import "fmt"

var picked = 6

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 0, n

	for l <= r {
		mid := (l + r) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func guess(value int) int {
	if value > picked {
		return -1
	}
	if value < picked {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(guessNumber(10))
}
