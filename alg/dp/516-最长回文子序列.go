package main

func longestPalindromeSubseq(s string) int {
	sLen := len(s)
	dp := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]int, sLen+1)
	}

	for i := sLen - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < sLen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][sLen-1]
}

//状态压缩

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
