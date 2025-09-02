package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i := 0
	for j := 0; j < len(s); j++ {
		if i < len(g) && s[j] >= g[i] {
			i++
		}
	}

	return i
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}
