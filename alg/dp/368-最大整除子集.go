package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	l := len(nums)
	max, end := 0, -1
	memo, last := make([]int, l), make([]int, l)
	for k := range memo {
		memo[k] = 1
		last[k] = -1
	}
	sort.Ints(nums)
	for right := 0; right < l; right++ {
		for left := 0; left < right; left++ {
			// 如果右边的数能整除左边的数字，且当前右边的数字构成的序列较短
			if nums[right]%nums[left] == 0 && memo[right] <= memo[left] {
				memo[right] = memo[left] + 1
				last[right] = left
			}
		}
		if memo[right] > max {
			max = memo[right]
			end = right
		}
	}
	var result []int
	//倒序输出
	for i := end; i != -1; i = last[i] {
		result = append([]int{nums[i]}, result...)
	}
	return result
}
