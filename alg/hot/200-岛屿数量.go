package hot

//DFS
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])

	num_islands := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if grid[r][c] == '1' {
				num_islands++
				dfs(grid, r, c)
			}
		}
	}
	return num_islands
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}



//方法二BFS


//方法三：并查集
type UnionFindSet struct {
	Parents []int // 每个结点的顶级节点
	SetCount int // 连通分量的个数
}

func (u *UnionFindSet) Init(grid [][]byte) {
	row := len(grid)
	col := len(grid[0])
	count := row*col
	u.Parents = make([]int, count)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			u.Parents[i*col+j] = i*col+j
			if grid[i][j] == '1' {
				u.SetCount++
			}
		}
	}
}

func (u *UnionFindSet) Find(node int) int {
	if u.Parents[node] == node {
		return node
	}
	root := u.Find(u.Parents[node])
	u.Parents[node] = root
	return root
}

func (u *UnionFindSet) Union(node1 int, node2 int) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)
	if root1 == root2 {
		return
	}
	if root1 < root2 {
		u.Parents[root1] = root2
	} else {
		u.Parents[root2] = root1
	}
	u.SetCount--
}

// 心得：并查集是一种搜索算法（针对聚合的）
func numIslandsV3(grid [][]byte) int {
	// 创建并初始化并查集
	u := &UnionFindSet{}
	row := len(grid)
	col := len(grid[0])
	u.Init(grid)
	// 根据grid建立相应的并查集，并统计连通分量个数【每连接一次进行减一】
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// 如果周边四个方向也是1就进行union
				if i - 1 >= 0 && grid[i-1][j] == '1' {
					u.Union(i*col+j, (i-1)*col+j)
				}
				if i + 1 < row && grid[i+1][j] == '1' {
					u.Union(i*col+j, (i+1)*col+j)
				}
				if j - 1 >= 0 && grid[i][j-1] == '1' {
					u.Union(i*col+j, i*col+(j-1))
				}
				if j + 1 < col && grid[i][j+1] == '1' {
					u.Union(i*col+j, i*col+(j+1))
				}
				grid[i][j] = '0'
			}
		}
	}
	// 返回结果
	return u.SetCount
}

