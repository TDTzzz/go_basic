package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	p := []*TreeNode{root}

	for i := 0; len(p) > 0; i++ {
		res = append(res, []int{})
		q := []*TreeNode{}
		for j := 0; j < len(p); j++ {
			node := p[j]
			res[i] = append(res[i], node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		p = q
	}
	return res
}
