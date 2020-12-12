package main

//动态规划
func rob(nums []int) (res int) {
	if len(nums) == 0 {
		return
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		//偷还是不偷
		stealRes := nums[i] + dp[i-2]
		noStealRes := dp[i-1]
		dp[i] = max(stealRes, noStealRes)
	}
	return dp[len(nums)-1]
}

//空间优化版动态规划
func robv2(nums []int) int {
	pre1, pre2 := 0, 0

	for _, v := range nums {
		cur := max(pre2+v, pre1)
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
