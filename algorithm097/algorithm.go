package algorithm097

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return len(s3) == 0
	}
	if len(s1) == 0 && len(s2) != 0 {
		return s2 == s3
	}
	if len(s1) != 0 && len(s2) == 0 {
		return s1 == s3
	}
	len1, len2 := len(s1), len(s2)
	if len(s3) == 0 || len1+len2 != len(s3) {
		return false
	}
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
		for j := range dp[i] {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1]
			} else if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1]
			} else {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] ||
					dp[i][j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[len1][len2]
}

// 超时
/*func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if s1[0] != s2[0] {
		if s3[0] != s1[0] && s3[0] != s2[0] {
			return false
		}
		if s3[0] == s1[0] {
			return isInterleave(s1[1:], s2, s3[1:])
		}
		if s3[0] == s2[0] {
			return isInterleave(s1, s2[1:], s3[1:])
		}
	}
	// s1[0] == s2[0]
	if s3[0] != s1[0] {
		return false
	}
	return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
}*/
