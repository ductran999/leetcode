package main

import "fmt"

func numberOfBeams(bank []string) int {
	ans := 0
	prev := 0
	for _, row := range bank {
		cnt := 0
		for i := 0; i < len(row); i++ {
			if row[i] == '1' {
				cnt++
			}
		}
		if cnt > 0 {
			ans += prev * cnt
			prev = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
}
