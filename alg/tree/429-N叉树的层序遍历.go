package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}

	for level := 0; len(queue) > 0; level++ {
		res = append(res, []int{})
		cnt := len(queue)
		for j := 0; j < cnt; j++ {
			if queue[j] != nil {
				res[level] = append(res[level], queue[j].Val)
				for _, n := range queue[j].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[cnt:]
	}
	return
}
