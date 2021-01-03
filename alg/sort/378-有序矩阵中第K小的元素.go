package main

import "log"

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	res := kthSmallest(matrix, 8)
	log.Println(res)
}

//二分法
func kthSmallest(matrix [][]int, k int) int {
	d := len(matrix)
	mid, lo, hi := 0, matrix[0][0], matrix[d-1][d-1]
	for lo <= hi {
		mid = (lo + hi) / 2
		curr := check(matrix, mid)
		if curr == k {
			if check(matrix, mid-1) < k {
				return mid
			}
			hi = mid - 1
		} else if curr > k {
			if check(matrix, mid-1) < k {
				return mid
			}
			hi = mid - 1
		} else if curr < k {
			lo = mid + 1
		}
	}
	return mid
}

func check(matrix [][]int, target int) int {
	d := len(matrix)
	count, i, j := 0, d-1, 0
	for i >= 0 && j < d {
		if matrix[i][j] <= target {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count
}
