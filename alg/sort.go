package main

import "log"

func main() {
	//res := []int{1, 2, 3}
	res := []int{4, 3, 5, 2, 1}
	res2 := selectTop2(res)

	log.Println(res2)
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

//On求数组第二大的数，用栈思想
func selectTop2(nums []int) int {
	var top, second int
	if nums[0] > nums[1] {
		top = nums[0]
		second = nums[1]
	} else {
		top = nums[1]
		second = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		curr := nums[i]
		if curr > top {
			second = top
			top = curr
			continue
		}

		if curr > second {
			second = curr
		}
	}
	return second
}

//全排列
func permutation(nums []int) [][]int {
	res := make([][]int, 0)
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, path)
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}

//归并排序
func merge(left []int, right []int) []int {
	i, j := 0, 0
	newArr := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			newArr = append(newArr, right[j])
			j++
		} else {
			newArr = append(newArr, left[i])
			i++
		}
	}
	//两边还有剩余
	if i < len(left) {
		newArr = append(newArr, left[i:]...)
	}
	if j < len(right) {
		newArr = append(newArr, left[j:]...)
	}
	return newArr
}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	left := MergeSort(arr[:key])
	right := MergeSort(arr[key:])
	return merge(left, right)
}

//
func test3(nums []int, dad int, end int) {
	son := 2*dad + 1
	if nums[son+1] > nums[son] {
		son++
	}

	if nums[dad] < nums[son] {
		nums[dad], nums[son] = nums[son], nums[dad]
	}

	test3(nums, son, end)
}
