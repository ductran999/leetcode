package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	place := 0

	if len(chars) <= 1 {
		return len(chars)
	}

	slow := 0
	for fast := 1; fast <= len(chars); fast++ {
		// detect last element of array or current char is different from the previous
		if fast == len(chars) || chars[fast] != chars[fast-1] {
			times := fast - slow

			ch := chars[slow]
			chars[place] = ch
			place++

			// number compress
			if times > 1 {
				digits := []byte(strconv.Itoa(times))
				for _, d := range digits {
					chars[place] = d
					place++
				}
			}

			// update pivot for new array
			slow = fast
		}
	}

	return place
}

func main() {
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars))
}
