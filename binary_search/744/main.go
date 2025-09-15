package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1

	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if letters[l] > target && l <= r {
		return letters[l]
	}
	return letters[0]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := 'c'
	fmt.Println(string(nextGreatestLetter(letters, byte(target))))
}
