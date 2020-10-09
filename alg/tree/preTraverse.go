package main

import (
	"go_basic/alg/structures"
)

//前序遍历-根左右
func preorderTraversal(root *structures.TreeNode) []int {
	var (
		res   []int
		stack []*structures.TreeNode
	)
	stack = append(stack, root)

	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		//出栈

	}

	return res
}

//中序遍历
func () {

}

//后序遍历
func () {

}
