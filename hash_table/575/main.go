package main

import "fmt"

func distributeCandies(candyType []int) int {
	// Alice can eat only half of the total candies
	n := len(candyType) / 2

	// Use a fixed-size array as a set to track distinct candy types
	// Constraint: -10^5 <= candyType[i] <= 10^5
	// So we need an array of size 200_001 (shifted by +100_000)
	offset := 100_000
	seen := make([]bool, 200_001)

	distinct := 0
	for _, candy := range candyType {
		idx := candy + offset
		if !seen[idx] {
			seen[idx] = true
			distinct++

			// Early exit: Alice already has enough unique candy types
			if distinct == n {
				return n
			}
		}
	}

	return distinct
}

func main() {
	candyType := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candyType))
}
