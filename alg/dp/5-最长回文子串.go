package main

//DP
func longestPalindrome(s string) string {
	sLen := len(s)
	dp := make([][]bool, sLen)

	maxLen := 1
	begin := 0
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
	}

	//子串长度
	for l := 2; l <= sLen; l++ {

		for i := 0; i < sLen; i++ {
			j := i + l - 1

			if j >= sLen {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

//中心扩展法
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
