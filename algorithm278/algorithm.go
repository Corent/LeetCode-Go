package algorithm278

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) >> 1
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
