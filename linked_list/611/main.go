package main

import "fmt"

type Info struct {
	freq    int
	prev    int
	numLess int
}

func triangleNumber(nums []int) int {
	const maxVal = 1000
	info := make([]Info, maxVal+1)
	for i := range info {
		info[i].numLess = len(nums)
	}
	for _, v := range nums {
		info[v].freq++
	}

	// find biggest number that appears
	biggest := maxVal
	for biggest > 0 && info[biggest].freq == 0 {
		biggest--
	}
	if biggest == 0 {
		return 0
	}

	// Build inplace linked list
	next := biggest
	info[biggest].numLess = len(nums) - (info[biggest].freq + info[0].freq)
	for current := biggest - 1; current > 0; current-- {
		data := &info[current]
		data.numLess = info[next].numLess - data.freq
		if data.freq > 0 {
			info[next].prev = current
			next = current
		}
	}

	r := 0
	for a := biggest; a > 0; a = info[a].prev {
		fa := info[a].freq

		// all three come from a
		r += (fa * (fa - 1) * (fa - 2)) / 6

		// two come from a and third from others
		if fa > 1 {
			r += (fa * (fa - 1) / 2) * info[a].numLess
		}

		badB := a / 2
		kb := 0
		for b := info[a].prev; b > badB; b = info[b].prev {
			fb := info[b].freq
			// a b b
			kb += (fb * (fb - 1)) / 2
			// abc
			kb += fb * (info[b].numLess - info[a+1-b].numLess)
		}
		r += kb * fa
	}

	return r
}

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
}
