package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	w := 0
	b := 0

	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			w++
		} else {
			b++
		}
	}
	if b == k {
		return 0
	}

	min := w
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			w++
		} else {
			b++
		}

		if blocks[i-k] == 'W' {
			w--
		} else {
			b--
		}

		if w < min {
			min = w
		}
	}

	return min
}

func main() {
	blocks := "WBBWWBBWBW"
	k := 7
	fmt.Println(minimumRecolors(blocks, k))
}
