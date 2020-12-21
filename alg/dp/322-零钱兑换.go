package main

//动态规划（压缩空间）
func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j > coins[i-1] {
				dp[j] = Min(dp[j], 1+dp[j-coins[i-1]])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
