package main

//思路1：动态规划
//思路2：Dijkstra
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	pv := make(map[int][][]int)
	for _, v := range flights {
		if _, ok := pv[v[0]]; !ok {
			pv[v[0]] = make([][]int, 0)
		}
		pv[v[0]] = append(pv[v[0]], []int{v[1], v[2]})
	}
	min, i := -1, -1
	if dfs, ok := pv[src]; ok {
		for {
			arr := dfs
			if len(arr) == 0 || i >= K {
				break
			}
			for _, v := range arr {
				if v[0] == dst {
					if v[1] < min || min == -1 {
						min = v[1]
					}
					continue
				}
				if pArr, ok := pv[v[0]]; ok {
					for _, tmp := range pArr {
						if tmp[1]+v[1] > min && min != -1 {
							continue
						}
						dfs = append(dfs, []int{tmp[0], tmp[1] + v[1]})
					}
				}
			}
			dfs = dfs[len(arr):]
			i++
		}
	}
	return min
}
