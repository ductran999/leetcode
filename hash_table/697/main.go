package main

import "fmt"

func findShortestSubArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	freq := make(map[int]int)
	max := 0
	for _, val := range nums {
		freq[val]++
		if freq[val] > max {
			max = freq[val]
		}
	}

	if max == 1 {
		return 1
	}

	reachMax := []int{}
	for k, val := range freq {
		if val == max {
			reachMax = append(reachMax, k)
		}
	}

	minLen := len(nums)
	for _, need := range reachMax {
		start := -1
		end := -1
		for i, val := range nums {
			if val == need {
				if start == -1 {
					start = i
				} else {
					end = i
				}
			}
		}

		if end-start+1 < minLen {
			minLen = end - start + 1
		}
	}

	return minLen
}

func main() {
	nums := []int{2, 1}
	fmt.Println(findShortestSubArray(nums))
}
