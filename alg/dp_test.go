package main

import (
	"math"
	"sort"
)

//动态规划

//最长上升自序列
func lis(nums []int) (res int) {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, 0)
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return
}

//信封嵌套
func maxEnvelopes(envelopes [][]int) int {
	n1 := len(envelopes)
	if n1 <= 1 {
		return n1
	}
	//对宽度进行排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	//剩下的就简单了，对高度进行LIS
	res := 0
	dp := make([]int, n1)
	for i := 0; i < n1; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

//最大子数组
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	res := math.MinInt16
	dp := make([]int, l)
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i], dp[i-1]+nums[i])
	}

	return res
}

//辅助函数
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
