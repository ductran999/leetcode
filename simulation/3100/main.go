package main

import (
	"fmt"
	"math"
)

func maxBottlesDrunk(numBottles int, numExchange int) int {
	b := float64(2*numExchange - 3)
	c := float64(-2*numBottles + 2)
	extra := int((-b + math.Sqrt(b*b-4*c)) / 2.0)

	return numBottles + extra
}

func main() {
	numBottles, numExchange := 13, 6
	fmt.Println(maxBottlesDrunk(numBottles, numExchange))
}
