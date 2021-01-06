package main

func pre(root *TreeNode) (res []int) {

}

func inorder(root *TreeNode) (res []int) {

}

func postorder(root *TreeNode) (res []int) {
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
	return
}
