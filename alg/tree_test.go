package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := preorderTraversal(getExample())
	log.Println(res)
}

func preorderTraversal(root *TreeNode) (res []int) {

	stk := make([]*TreeNode, 0)

	for len(stk) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stk = append(stk, root)
			root = root.Left
		}
		index := len(stk) - 1
		root = stk[index].Right
		stk = stk[:index]
	}
	return
}

func postorderTraversal(root *TreeNode) (res []int) {
	stk := make([]*TreeNode, 0)

	var prev *TreeNode
	for len(stk) > 0 || root != nil {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		index := len(stk) - 1
		root = stk[index]
		stk = stk[:index]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root.Right)
			root = root.Right
		}
	}
	return res
}

func getExample() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
}

//层序遍历
func levelOrder(root *TreeNode) (res [][]int) {
	queue := make([]*TreeNode, 0)
	level := 0
	queue = append(queue, root)

	for len(queue) > 0 {
		res = append(res, []int{})
		tmpQueue := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
		level++
	}
	return res
}

//lca问题
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}

//dp
func coinChange(coins []int, amount int) (res int) {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i-1] < j {
				dp[j] = min(dp[j], dp[j-coins[i-1]]+1)
			}
		}
	}
	return
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
