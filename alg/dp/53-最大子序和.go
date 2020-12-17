package main

//dp
func maxSubArray(nums []int) int {
	n1 := len(nums)
	dp := make([]int, n1)

	for i := 0; i < n1; i++ {
		dp[i] = nums[i]
	}
	max := dp[0]
	for i := 1; i < n1; i++ {
		dp[i] = Max(dp[i], nums[i]+dp[i-1])
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

//分治

func maxSubArrayV2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := Max(l.lSum, l.iSum+r.lSum)
	rSum := Max(r.rSum, r.iSum+l.rSum)
	mSum := Max(Max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
