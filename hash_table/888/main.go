package main

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	ans := make([]int, 2)

	aCandies := make(map[int]int)
	bCandies := make(map[int]int)

	aTotal := 0
	for _, val := range aliceSizes {
		aCandies[val]++
		aTotal += val
	}

	bTotal := 0
	for _, val := range bobSizes {
		bCandies[val]++
		bTotal += val
	}

	for aCandy := range aCandies {
		for bCandy := range bCandies {
			aliceTotalExchanged := aTotal - aCandy + bCandy
			bobTotalExchanged := bTotal - bCandy + aCandy
			if aliceTotalExchanged == bobTotalExchanged {
				ans[0] = aCandy
				ans[1] = bCandy
				return ans
			}
		}
	}

	return ans
}

func main() {
	aliceSizes := []int{1, 1}
	bobSizes := []int{2, 2}
	fmt.Println(fairCandySwap(aliceSizes, bobSizes))
}
