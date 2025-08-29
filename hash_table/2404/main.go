package main

import "fmt"

func mostFrequentEven(nums []int) int {
	freq := make(map[int]int)

	for _, val := range nums {
		if val%2 == 0 {
			freq[val]++
		}
	}
	if len(freq) == 0 {
		return -1
	}

	count := 0
	res := -1
	for k, v := range freq {
		if v < count {
			continue
		}

		if res == -1 || v > count {
			count = v
			res = k
			continue
		}

		if v == count && k < res {
			res = k
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	fmt.Println(mostFrequentEven(nums))
}
