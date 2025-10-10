package main

import "fmt"

func maximumEnergy(energy []int, k int) int {
	max := -1000

	for i := len(energy) - 1; i >= 0; i-- {
		if i+k < len(energy) {
			energy[i] += energy[i+k]
		}
		if energy[i] > max {
			max = energy[i]
		}
	}
	return max
}

func main() {
	energy := []int{-2, -3, -1}
	k := 2
	fmt.Println(maximumEnergy(energy, k))
}
