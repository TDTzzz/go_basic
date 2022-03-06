package main

func numIslands(grid [][]byte) (res int) {

	r := len(grid)
	c := len(grid[0])

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {

			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return
}

func dfs(grid [][]byte, r, c int) {

	if r <= 0 || c <= 0 || r >= len(grid) || c >= len(grid[0]) {
		return
	}
	grid[r][c] = '1'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
