package main

import "log"

func main() {
	s := "aa"
	p := "*"
	res := isMatch(s, p)
	log.Println(res)
}

//递归的DP没搞出来
//func isMatch(s string, p string) bool {
//	var dp func(i, j int) bool
//	dp = func(i, j int) bool {
//		if i+1 > len(s) {
//			return false
//		}
//		if j == len(p) {
//			return i == len(s)
//		}
//		if s[i] == p[j] || string(p[j]) == "?" {
//			//判断有没有通配符
//			if j+1 < len(p) && string(p[j+1]) == "*" {
//				return dp(i, j+2) || dp(i+1, j)
//			} else {
//				return dp(i+1, j+1)
//			}
//		} else {
//			if string(p[0]) == "*" {
//				return dp(i+1, j)
//			}
//			if j+1 < len(p) && string(p[j+1]) == "*" {
//				return dp(i, j+2)
//			} else {
//				return false
//			}
//		}
//	}
//	return dp(0, 0)
//}

//dp table
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
