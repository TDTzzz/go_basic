package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DFS
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		res += dfs(root.Right)
	}
	if root.Left != nil {
		if isLeafNode(root.Left) {
			res += root.Left.Val
		} else {
			res += dfs(root.Left)
		}
	}
	return
}

//判断是否是叶子节点
func isLeafNode(root *TreeNode) bool {
	return root.Right == nil && root.Left == nil
}

//BFS
func sumOfLeftLeavesV2(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if isLeafNode(node.Left) {
				res += node.Left.Val
			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return
}
