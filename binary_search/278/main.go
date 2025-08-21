package main

var badVer = 4

func isBadVersion(version int) bool {
	return true
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if isBadVersion(mid) { // mid match the condition narrow
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

}
