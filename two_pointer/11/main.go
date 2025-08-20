package main

import "fmt"

func calcMax(currMax, h, w int) int {
	area := h * w
	if area > currMax {
		return area
	}
	return currMax
}

func maxArea(height []int) int {
	res := -1
	l, r := 0, len(height)-1

	for l < r {
		width := r - l
		switch {
		case height[l] < height[r]:
			res = calcMax(res, height[l], width)
			l++
		case height[l] > height[r]:
			res = calcMax(res, height[r], width)
			r--
		default:
			res = calcMax(res, height[r], width)
			l++
			r--
		}
	}

	return res
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
