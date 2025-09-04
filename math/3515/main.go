package main

import "fmt"

func distance(target, start int) int {
	if start > target {
		return start - target
	}
	return target - start
}

func findClosest(x int, y int, z int) int {
	p1 := distance(x, z)
	p2 := distance(y, z)

	if p1 > p2 {
		return 2
	}
	if p2 > p1 {
		return 1
	}

	return 0
}

func main() {
	x := 2
	y := 7
	z := 4

	fmt.Println(findClosest(x, y, z))
}
