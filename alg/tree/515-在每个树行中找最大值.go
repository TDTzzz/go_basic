package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for level := 0; len(queue) > 0; level++ {

		length := len(queue)
		max := queue[0].Val
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if max < node.Val {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, max)
	}
	return
}
