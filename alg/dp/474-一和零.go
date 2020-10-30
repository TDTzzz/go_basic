package main

import "log"

//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
//请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
//如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

func findMaxForm(strs []string, m int, n int) int {
	statistic := func(str string) (int, int) {
		var zero, one = 0, 0
		for _, char := range str {
			if char == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}
	dp := make([][]int, m+1)
	//初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//进行动态规划
	for _, str := range strs {
		zero, one := statistic(str)
		//当zero或者one比m或者n大时，这个字符串就不用判断了。
		if zero > m || one > n {
			continue
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	res := findMaxForm(strs, m, n)
	log.Println(res)
}
