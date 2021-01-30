package main

//模拟

//按层模拟
func spiralOrder(matrix [][]int) []int {
	w, h := len(matrix[0]), len(matrix)

	left, right := 0, w-1
	low, height := 0, h-1

	res := make([]int, 0)
	for left <= right && low <= height {
		for column := left; column <= right; column++ {
			res = append(res, matrix[low][column])
		}
		for row := low + 1; row <= height; row++ {
			res = append(res, matrix[row][right])
		}
		if left < right && low < height {
			for column := right - 1; column > left; column-- {
				res = append(res, matrix[height][column])
			}

			for row := height; row > low; row-- {
				res = append(res, matrix[row][left])
			}
		}

		left++
		right--
		low++
		height--
	}
	return res
}
