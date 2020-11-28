package main

import (
	"log"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := 0
	depth(root, &m)
	return m
}

func depth(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, m)
	right := depth(root.Right, m)
	//此节点的直径和现有最长直径比较
	*m = int(max(left+right, *m))
	return int(max(left, right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}

	res := diameterOfBinaryTree(root)
	log.Println(res)
}
