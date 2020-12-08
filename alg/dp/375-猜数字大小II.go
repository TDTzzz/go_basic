package main

import "math"

//普通DP
func getMoneyAmount(n int) int {
	memo := make([][]int, n+1)
	for k := range memo {
		memo[k] = make([]int, n+1)
	}
	for l := 2; l <= n; l++ {
		for left := 1; left+l-1 <= n; left++ {
			minValue := math.MaxInt32
			right := left + l - 1
			for piv := left; piv < right; piv++ {
				minValue = min(minValue, piv+max(memo[left][piv-1], memo[piv+1][right]))
			}
			memo[left][right] = minValue
		}
	}
	return memo[1][n]
}

//优化DP
func getMoneyAmountV2(n int) int {
	dp := make([][]int, n+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for len := 2; len <= n; len++ {
		for left := 1; left <= n-len+1; left++ {
			minres := math.MaxInt32
			for piv := left + (len-1)/2; piv < left+len-1; piv++ {
				res := piv + max(dp[left][piv-1], dp[piv+1][left+len-1])
				minres = min(res, minres)
			}
			dp[left][left+len-1] = minres
		}
	}
	return dp[1][n]
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
