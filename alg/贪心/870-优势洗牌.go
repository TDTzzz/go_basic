package 贪心

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	pos := [][2]int{}
	for index, value := range nums2 {
		pos = append(pos, [2]int{value, index})
	}

	sort.Ints(nums1)
	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0] < pos[j][0]
	})
	resp := make([]int, len(nums1))

	temp := 0
	for i := len(pos) - 1; i >= 0; i-- {

		if nums1[i+temp] > pos[i][0] {
			resp[pos[i][1]] = nums1[i+temp]
		} else {
			resp[pos[i][1]] = nums1[temp]
			temp++
		}

	}

	return resp
}
