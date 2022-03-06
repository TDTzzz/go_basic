package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//BFS
func connect(root *Node) *Node {

	if root == nil {
		return root
	}
	stack := make([]*Node, 0)

	stack = append(stack, root)

	for len(stack) > 0 {

		levelCnt := len(stack)
		var pre *Node
		for i := 0; i < levelCnt; i++ {
			node := stack[0]
			stack = stack[1:len(stack)]
			if pre != nil {
				pre.Next = node
			}
			pre = node
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return root
}
