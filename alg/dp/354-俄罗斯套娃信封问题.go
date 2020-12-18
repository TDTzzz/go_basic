package main

import "sort"

//思路:
//1.对宽度进行升序排序（宽度一样的，高度倒序）
//2.问题被转化成了LIS（最长递增子序列）
func maxEnvelopes(envelopes [][]int) int {
	//按宽度排序
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
	//排序完了，此题就相当于是LIS（最长递增子序列）(dp/2分)
	res := 0
	dp := make([]int, n1)
	for i := 0; i < n1; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
