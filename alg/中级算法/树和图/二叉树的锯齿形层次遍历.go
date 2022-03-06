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

func zigzagLevelOrder(root *TreeNode) (res [][]int) {

	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for i := 0; len(stack) > 0; i++ {

		nextStack := make([]*TreeNode, 0)
		res = append(res, []int{})
		for j := 0; j < len(stack); j++ {

			curRoot := stack[j]
			if curRoot.Left != nil {
				nextStack = append(nextStack, curRoot.Left)
			}

			if curRoot.Right != nil {
				nextStack = append(nextStack, curRoot.Right)
			}

			res[i] = append(res[i], curRoot.Val)
		}
		//奇数层反转
		if i%2 == 1 {
			for a, n := 0, len(res[i]); a < n/2; a++ {
				res[i][a], res[i][n-1-a] = res[i][n-1-a], res[i][a]
			}
		}
		stack = nextStack
	}

	return
}

//bfs广度优先遍历
