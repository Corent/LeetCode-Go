package algorithm087

// https://blog.csdn.net/qq_46105170/article/details/108577730

var (
	str1, str2 string
	memo       [][][]int // memo[i][j][n]表示 s1[i: i + n] 和 s2[j: j + n] 是否可达，0: 未计算 1: 可达 -1: 不可达
)

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	n := len(s1)
	str1, str2, memo = s1, s2, make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, n+1)
		}
	}
	return isScrambleCore(0, 0, n)
}

func isScrambleCore(idx1, idx2, n int) bool {
	if memo[idx1][idx2][n] != 0 {
		return memo[idx1][idx2][n] > 0
	}
	if str1[idx1:idx1+n] == str2[idx2:idx2+n] {
		memo[idx1][idx2][n] = 1
		return true
	}
	hash := make([]uint8, 26)
	for i := idx1; i < idx1+n; i++ {
		hash[str1[i]-'a']++
	}
	for i := idx2; i < idx2+n; i++ {
		if hash[str2[i]-'a'] == 0 {
			memo[idx1][idx2][n] = -1
			return false
		}
		hash[str2[i]-'a']--
	}
	for i := 1; i < n; i++ {
		if isScrambleCore(idx1, idx2, i) && isScrambleCore(idx1+i, idx2+i, n-i) ||
			isScrambleCore(idx1, idx2+n-i, i) && isScrambleCore(idx1+i, idx2, n-i) {
			memo[idx1][idx2][n] = 1
			return true
		}
	}
	memo[idx1][idx2][n] = 0
	return false
}

// 超时，但Java版能过
/*func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	hash := [26]uint8{}
	for i := range s1 {
		hash[s1[i]-'a']++
		hash[s2[i]-'a']--
	}
	for i := range hash {
		if hash[i] != 0 {
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) ||
			isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}*/
