package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}

func main() {
	fmt.Println(numWaterBottles(10, 3))
}
