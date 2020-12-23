package main

import "log"

func main() {
	res := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	//res := []int{0, 1, 0, 3, 2, 3}
	//res := []int{1, 2, 8, 6, 4}

	nn := LISV2(res)
	log.Println(nn)
}

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

//返回数组（牛客版本还没通过）
func LIS(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	if length == 0 {
		return res
	}
	dp, maxNum := make([]int, length), 0

	tmp := make([]int, 0)
	for i := 0; i < length; i++ {
		tmp = []int{}
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

				if dp[i] < dp[j]+1 {
					//确保是最小的
					if nums[j+1] < nums[j] && dp[i] < dp[j+1]+1 {
						continue
					}

					dp[i] = dp[j] + 1
					tmp = append(tmp, nums[j])
				}
			}
		}
		tmp = append(tmp, nums[i])
		if maxNum <= dp[i] {
			maxNum = dp[i]
			res = tmp
		}
	}
	return res
}

func LISV2(nums []int) []int {
	length := len(nums)
	res := make([]int, 0)
	maxLen := make([]int, 0)

	for i := 0; i < length; i++ {
		if i == 0 {
			res = append(res, nums[i])
			maxLen = append(maxLen, len(res))
		} else {
			if res[len(res)-1] < nums[i] {
				res = append(res, nums[i])
				maxLen = append(maxLen, len(res))
			} else {
				for j := 0; j < len(res); j++ {
					if res[j] > nums[i] {
						res[j] = nums[i]
						maxLen = append(maxLen, j+1)
						break
					}
				}
			}

		}
	}

	resArray := make([]int, len(res))

	//从后往前遍历
	j := len(res)
	for i := length - 1; j > 0; i-- {
		if maxLen[i] == j {
			j -= 1
			resArray[j] = nums[i]
		}
	}
	return resArray
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
