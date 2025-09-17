package main

import (
	"fmt"
	"math"
)

func maximumPopulation(logs [][]int) int {
	yearSet := make(map[int]int)
	ans := math.MaxInt
	for _, v := range logs {
		yearSet[v[0]] = 0
		yearSet[v[1]] = 0
	}

	for k := range yearSet {
		for _, v := range logs {
			if k >= v[0] && k <= v[1]-1 {
				yearSet[k]++
			}
		}
	}

	maxPopulation := 0
	for year, population := range yearSet {
		if population > maxPopulation {
			maxPopulation = population
			ans = year
		} else if population == maxPopulation && year < ans {
			ans = year
		}
	}

	return ans
}

func main() {
	logs := [][]int{{1993, 1999}, {2000, 2010}}
	fmt.Println(maximumPopulation(logs))
}
