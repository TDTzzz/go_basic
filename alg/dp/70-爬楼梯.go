package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 2, 1
	for i := 2; i < n; i++ {
		cur := pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}
