package main

import (
	"go_basic/alg/structures"
	"log"
)

func main() {
	res := preTraversalV3(mockData())
	log.Println(res)
}

//模拟数据
func mockData() *structures.TreeNode {

	root := structures.TreeNode{
		Val: 1,
		Right: &structures.TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
		Left: &structures.TreeNode{
			Val:   2,
			Right: nil,
			Left:  nil,
		},
	}

	return &root
}

//前序遍历-根左右
//V1-迭代法 时间复杂度On 空间复杂度 On
func preTraversalV1(root *structures.TreeNode) []int {
	var (
		stack []*structures.TreeNode
		res   []int
	)

	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:(len(stack) - 1)]
		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

//V2-莫里斯遍历(可以优化空间复杂度)
func preTraversalV2(root *structures.TreeNode) []int {
	var res []int
	node := root
	for node != nil {
		if node.Left != nil {
			res = append(res, node.Val)
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				res = append(res, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				node = node.Right
			}
		}
	}

	return res
}

//V3-迭代优化-只压右节点
func preTraversalV3(root *structures.TreeNode) []int {
	var (
		stack []*structures.TreeNode
		res   []int
	)
	if root == nil {
		return res
	}
	stack = append(stack, root)
	node := root

	for {
		if node != nil {
			res = append(res, node.Val)

			//只用压右节点
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			node = node.Left
		} else if len(stack) == 0 {
			return res
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
