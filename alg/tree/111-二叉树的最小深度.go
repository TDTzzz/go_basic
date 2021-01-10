package main

import (
	"log"
	"math"
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

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := minDepthV2(root)
	log.Println(res)
}

//DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

//BFS
func minDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for i := 1; len(queue) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			curr := queue[j]
			if curr.Left == nil && curr.Right == nil {
				return i
			}
			if curr.Left != nil {
				p = append(p, curr.Left)
			}
			if curr.Right != nil {
				p = append(p, curr.Right)
			}
		}
		queue = p
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
