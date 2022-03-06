//https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvmy42/

package main

func setZeroes(matrix [][]int) {
	//1.第一行第一列作为临时数组

	high := len(matrix)
	width := len(matrix[0])

	row := false
	col := false

	for i := 0; i < high; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					row = true
				}
				if j == 0 {
					col = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	//把那些应该为0的行和列全部置为0
	for i := 1; i < high; i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if col {
		for i := 0; i < high; i++ {
			matrix[i][0] = 0
		}
	}

	if row {
		for j := 0; j < width; j++ {
			matrix[0][j] = 0
		}
	}

}
