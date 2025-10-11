package main

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	count := map[int]int{}
	for _, p := range power {
		count[p]++
	}

	// Extract and sort unique power levels
	unique := make([]int, 0, len(count))
	for p := range count {
		unique = append(unique, p)
	}
	sort.Ints(unique)

	n := len(unique)
	dp := make([]int64, n)
	dp[0] = int64(unique[0] * count[unique[0]])

	for i := 1; i < n; i++ {
		cur := int64(unique[i] * count[unique[i]])
		dp[i] = dp[i-1] // skip current

		// find previous non-conflicting index
		j := i - 1
		for j >= 0 && unique[i]-unique[j] <= 2 {
			j--
		}

		if j >= 0 {
			cur += dp[j]
		}

		if cur > dp[i] {
			dp[i] = cur
		}
	}

	return dp[n-1]
}

func main() {
	power := []int{1, 1, 3, 4}
	fmt.Println(maximumTotalDamage(power))
}
