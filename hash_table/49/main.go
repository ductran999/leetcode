package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	res := [][]string{}
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
