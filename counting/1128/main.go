package main

import "fmt"

func totalPairs(n int) int {
	return n * (n + 1) / 2
}

func sortDomino(domino []int) [2]int {
	if domino[0] > domino[1] {
		domino[0], domino[1] = domino[1], domino[0]
	}
	return [2]int{domino[0], domino[1]}
}

func numEquivDominoPairs(dominoes [][]int) int {
	freq := make(map[[2]int]int)
	for _, do := range dominoes {
		d := sortDomino(do)
		freq[d]++
	}
	res := 0
	for _, val := range freq {
		if val > 1 {
			res += totalPairs(val - 1)
		}
	}

	return res
}

func main() {
	dominoes := [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}
	fmt.Println(numEquivDominoPairs(dominoes))
}
