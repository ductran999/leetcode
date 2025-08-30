package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	if len(nums) == 0 {
		return res
	}

	l := 0
	for r := 1; r <= len(nums); r++ {
		// Reach end of array or gap
		if r == len(nums) || nums[r]-nums[r-1] > 1 {
			if l == r-1 {
				res = append(res, fmt.Sprintf("%d", nums[l]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
			}
			l = r
		}
	}

	return res
}

func main() {
	nums := []int{1, 3}
	fmt.Println(summaryRanges(nums))
}
