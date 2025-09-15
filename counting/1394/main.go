package main

import "fmt"

func findLucky(arr []int) int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}

	for k, v := range freq {
		if k == v {
			return k
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(findLucky(arr))
}
