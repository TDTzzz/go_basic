package main

import "log"

//非最优解
func lengthOfLISV1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp, res := make([]int, length), 0

	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLISV1(nums)
	log.Println(res)
}
