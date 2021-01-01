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

//LCS 最长公共子序列
func lcs(s, t []int) int {
	l := len(s)
	w := len(t)

	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= w; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l][w]
}

//lcs状态压缩
func lcsV2(s, t []int) int {
	dp := make([]int, len(t)+1)

	for i := 1; i <= len(s); i++ {
		last := 0 //j=1时重置"左上角"的值
		for j := 1; j <= len(t); j++ {
			tmp := dp[j] //相当于dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[j] = last + 1
			} else {
				dp[j] = max(tmp, dp[j-1])
			}
			last = tmp
		}
	}
	return dp[len(t)]
}

//最长回文子序列
func lps(s string) int {
	l := len(s)
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, l+1)
	}

	//反着遍历
	for i := l - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	//试着正着遍历一下
	//dp[0][0] = 1
	//for i := 1; i < l; i++ {
	//	dp[i][i] = 1
	//	for j := i - 1; j >= 0; j-- {
	//		if s[i] == s[j] {
	//			dp[j][i] = dp[j+1][i-1] + 2
	//		} else {
	//			dp[j][i] = max(dp[j][j-1], dp[j+1][i])
	//		}
	//	}
	//}
	return dp[0][l-1]
}

//以最小插入次数构造回文串
func minInsertions(s string) int {
	l := len(s)
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, l+1)
	}

	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][l-1]
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
