package algorithm374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	l, r := 1, n
	for l < r {
		m := (l + r) >> 1
		if c := guess(m); c == 0 {
			return m
		} else if c < 0 {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
