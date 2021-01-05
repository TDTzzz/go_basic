package main

import "log"

func main() {
	//res := []int{1, 2, 3, 4, 5, 6}
	res := []int{6, 5, 4, 3, 2, 1}
	//res = quickSort(res)
	res = heapSort(res)
	log.Println(res)
}

func quickSort(nums []int) []int {
	left := 0
	right := len(nums)
	var help func(l, r int)
	help = func(l, r int) {
		if l < r {
			pivot := nums[(left+right)/2]
			i, j := left, right
			for {
				for nums[i] < pivot {
					i++
				}
				for nums[j] > pivot {
					j--
				}
				if i >= j {
					break
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
			help(l, i-1)
			help(j+1, r)
		}
	}
	return nums
}

func heapSort(nums []int) []int {
	//先构造最大堆
	for i := len(nums) / 2; i >= 0; i-- {
		heap_help(nums, i, len(nums)-1)
	}

	//取数
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heap_help(nums, 0, i-1)
	}
	return nums
}

func heap_help(nums []int, dad int, end int) {
	son := 2*dad + 1
	if son > end {
		return
	}
	if son+1 <= end && nums[son+1] > nums[son] {
		son++
	}

	if nums[dad] < nums[son] {
		nums[dad], nums[son] = nums[son], nums[dad]
	}
	heap_help(nums, son, end)
}
