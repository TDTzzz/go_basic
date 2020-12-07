package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//V1:DFS
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var res int

	for _, v := range root.Children {
		tmp := maxDepth(v)
		if tmp > res {
			res = tmp
		}
	}
	return 1 + res
}

//V2:BFS
func maxDepthV2(root *Node) int {
	if root == nil {
		return 0
	}
	level := 0
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level++
		l := len(queue)
		for i := 0; i < l; i++ {
			for _, v := range queue[i].Children {
				queue = append(queue, v)
			}
		}
		queue = queue[l:]
	}
	return level
}
