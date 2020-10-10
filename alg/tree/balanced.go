package main

import (
	"go_basic/alg"
	"go_basic/alg/structures"
)

//自顶向上 时间复杂度 On2 空间复杂度On
func isBalancedV2(root *structures.TreeNode) bool {
	if root == nil {
		return false
	}
	return alg.Abs(heightV2(root.Left)-heightV2(root.Right)) <= 1 && isBalancedV2(root.Left) && isBalancedV2(root.Right)
}

func heightV2(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	return alg.Max(heightV2(root.Left), heightV2(root.Right))
}

//自底向上 时间复杂度On 空间On
func isBalanced(root *structures.TreeNode) bool {
	return height(root) != -1
}

func height(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if alg.Abs(leftHeight-rightHeight) > 1 || leftHeight == -1 || rightHeight == -1 {
		return -1
	}
	return alg.Max(leftHeight, rightHeight) + 1
}
