package main

import "strings"

func main() {
	//res := solveNQueens(8)
	//log.Println(res)
	s := "...QQ.Q"
	//for k, v := range s {
	//	log.Println(k, string(v))
	//}

	s = s + "a"
}

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	res := [][]string{}
	helper(0, bd, &res, n)
	return res
}
func helper(r int, bd [][]string, res *[][]string, n int) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if isValid(r, c, n, bd) {
			bd[r][c] = "Q"
			helper(r+1, bd, res, n)
			bd[r][c] = "."
		}
	}

}
func isValid(r, c int, n int, bd [][]string) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}

//优化s
func solveNQueensV2(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	res := [][]string{}
	helperV2(0, bd, &res, n, &cols, &diag1, &diag2)
	return res
}

func helperV2(r int, bd [][]string, res *[][]string, n int, cols, diag1, diag2 *map[int]bool) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if !(*cols)[c] && !(*diag1)[r+c] && !(*diag2)[r-c] {
			bd[r][c] = "Q"
			(*cols)[c] = true
			(*diag1)[r+c] = true
			(*diag2)[r-c] = true
			helperV2(r+1, bd, res, n, cols, diag1, diag2)
			bd[r][c] = "."
			(*cols)[c] = false
			(*diag1)[r+c] = false
			(*diag2)[r-c] = false
		}
	}
}
