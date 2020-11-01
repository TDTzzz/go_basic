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

//贪心+二分
func lengthOfLISV2(nums []int) int {
	len, n := 1, len(nums)
	if n == 0 {
		return 0
	}
	d := make([]int, n+1)
	d[len] = nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > d[len] {
			len++
			d[len] = nums[i]
		} else {
			l, r, pos := 1, len, 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return len
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLISV2(nums)
	log.Println(res)
}
