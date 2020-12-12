package main

import "log"

//01背包
func knapsack01(v, w []int, capacity int) int {
	dp := make([][]int, len(v)+1)
	for i := 0; i <= len(v); i++ {
		dp[i] = make([]int, capacity+1)
	}

	//选第几个物品
	for i := 1; i <= len(v); i++ {
		for c := 0; c <= capacity; c++ {
			if c < w[i-1] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-w[i-1]]+v[i-1])
			}
		}
	}
	return dp[len(v)][capacity]
}

//空间优化版本
func knapsack01V2(v, w []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i := 1; i <= len(v); i++ {
		//for c := capacity; c >= 0; c-- {
		for c := 0; c <= capacity; c++ {
			if c >= w[i-1] {
				dp[c] = max(dp[c], dp[c-w[i-1]]+v[i-1])
			}
		}
		log.Println(dp)
	}
	return dp[capacity]
}

func main() {
	w := []int{1, 3, 4, 6, 8}
	v := []int{1, 7, 23, 16, 32}
	res := knapsack01(v, w, 12)
	res2 := knapsack01V2(v, w, 12)
	log.Println(res, res2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
