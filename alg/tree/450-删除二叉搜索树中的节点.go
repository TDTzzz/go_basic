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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//利用中序遍历的思想，递归 前驱节点 后驱节点进行节点的替换操作
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//到这里意味已经查找到目标
	if root.Right == nil {
		//右子树为空
		return root.Left
	}
	if root.Left == nil {
		//左子树为空
		return root.Right
	}
	//在右树中找后驱节点，左树中找前驱节点
	//minNode := root.Right
	//for minNode.Left != nil {
	//	//查找后继
	//	minNode = minNode.Left
	//}
	//root.Val = minNode.Val
	//root.Right = deleteMinNode(root.Right)

	ancestor := root
	successor := root.Right

	for successor.Left != nil {
		ancestor = successor
		successor = successor.Left
	}
	root.Val = successor.Val

	if successor == ancestor.Right {
		ancestor.Right = deleteNode(successor, successor.Val)
	} else {
		ancestor.Left = deleteNode(successor, successor.Val)
	}

	return root
}

//如果被删除节点是 leaf, 直接删除
//
//如果被删除节点 只有一个child, 使用仅有的 child 代替原节点
//
//如果被删除节点 有两个children, 则在 right subtree中 寻找 successor, 将原节点值替换为 successor 的值, 并递归删除 successor,

//func deleteMinNode(root *TreeNode) *TreeNode {
//	if root.Left == nil {
//		pRight := root.Right
//		root.Right = nil
//		return pRight
//	}
//	root.Left = deleteMinNode(root.Left)
//	return root
//}
