package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sFreq := make(map[byte]byte)
	tFreq := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		cs, ct := s[i], t[i]

		if val, ok := sFreq[cs]; ok {
			if val != ct {
				return false
			}
		} else {
			sFreq[cs] = ct
		}

		if val, ok := tFreq[ct]; ok {
			if val != cs {
				return false
			}
		} else {
			tFreq[ct] = cs
		}
	}

	return true
}

func main() {
	s := "foo"
	t := "egg"

	fmt.Println(isIsomorphic(s, t))
}
