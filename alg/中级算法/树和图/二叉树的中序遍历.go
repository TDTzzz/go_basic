package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

//递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {

		if root != nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

//迭代
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}

	return res
}
