package main

func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for k, _ := range dp {
		dp[k] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

//状态压缩
func longestCommonSubsequenceV2(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		last := 0 //左上的值
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] //即二维dp的dp[i-1][j]
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1
			} else {
				dp[j] = max(tmp, dp[j-1])
			}
			last = tmp
		}
	}
	return dp[len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


