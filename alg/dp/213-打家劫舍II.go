package main

import "log"

//其实就是在之前的基础上加了个 "偷不偷第一家"的判断
func rob(nums []int) int {
	pre1, pre2 := 0, 0
	pre11, pre22 := 0, 0

	if len(nums) == 1 {
		return nums[0]
	}
	//偷不偷第一个
	for i := 0; i < len(nums)-1; i++ {
		cur := max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}
	for i := 1; i < len(nums); i++ {
		curr := max(pre22+nums[i], pre11)
		pre22 = pre11
		pre11 = curr
	}
	log.Println(pre1, pre11)
	return max(pre1, pre11)
}

//把重复代码提出来的优化
func robV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelp(nums, 0, len(nums)-1), robHelp(nums, 1, len(nums)))
}

func robHelp(nums []int, first, last int) int {
	pre1, pre2 := 0, 0
	for i := first; i < last; i++ {
		cur := max(pre1, pre2+nums[i])
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
