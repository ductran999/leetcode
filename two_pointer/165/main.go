package main

import "fmt"

// parseRevision extracts the next numeric revision from version starting at index.
// It returns the revision number and the next index to continue parsing from.
func parseRevision(version string, index int) (int, int) {
	num := 0
	for index < len(version) && version[index] != '.' {
		num = num*10 + int(version[index]-'0')
		index++
	}
	// Skip the '.' if present
	if index < len(version) && version[index] == '.' {
		index++
	}
	return num, index
}

func compareVersion(version1, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		var rev1, rev2 int
		rev1, i = parseRevision(version1, i)
		rev2, j = parseRevision(version2, j)

		if rev1 < rev2 {
			return -1
		} else if rev1 > rev2 {
			return 1
		}
	}
	return 0
}

func main() {
	version1 := "1.01"
	version2 := "1.001"
	fmt.Println(compareVersion(version1, version2))
}
