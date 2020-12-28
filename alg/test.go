package alg

//练手

func main() {

}

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

func preorderTraversal(root *TreeNode) (res []int) {
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || root != nil {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func inorderTraversal(root *TreeNode) (res []int) {
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || root != nil {

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
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
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root.Right)
			root = root.Right
		}
	}
	return
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {

		queue_tmp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {

			res[i] = append(res[i], queue[j].Val)
			if queue[j].Left != nil {
				queue_tmp = append(queue_tmp, queue[j].Left)
			}

			if queue[j].Right != nil {
				queue_tmp = append(queue_tmp, queue[j].Right)
			}
		}
		queue = queue_tmp
	}

	return
}
