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
func wiggleMaxLengthV2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	begin, up, down := 0, 1, 2
	state := begin
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		switch state {
		case begin:
			if nums[i] > nums[i-1] {
				state = up
				maxLen++
			} else if nums[i] < nums[i-1] {
				state = down
				maxLen++
			}
		case up:
			if nums[i] < nums[i-1] {
				state = down
				maxLen++
			}
		case down:
			if nums[i] > nums[i-1] {
				state = up
				maxLen++
			}
		}
	}

	return maxLen
}
