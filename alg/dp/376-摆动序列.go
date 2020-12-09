package main

//空间优化后的动态规划
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	down, up := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if down > up {
		return down
	}
	return up
}

//贪心
func wiggleMaxLengthV3(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	prevdiff := nums[1] - nums[0]
	count := 2
	if prevdiff == 0 {
		count = 1
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if ((diff > 0) && prevdiff <= 0) || (diff < 0 && prevdiff >= 0) {
			count++
			prevdiff = diff
		}
	}
	return count
}
