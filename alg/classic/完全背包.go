package main

func knapsack(coins []int, amount int) int {
	coinsNum := len(coins)
	dp := make([][]int, coinsNum+1)
	for i := 0; i < coinsNum+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i < coinsNum+1; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[coinsNum][amount]
}

//状态压缩
func knapsackV2(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j <= amount; j++ {
			if j > coins[i] {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
