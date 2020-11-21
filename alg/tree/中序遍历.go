package main

//迭代
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}

	return res
}

//递归
func inorderTraversalV2(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

//最好的morris 空间复杂度为O1
//核心思想一句话:假设当前遍历到节点X，将x的前驱节点的右孩子指向x
func inorderTraversalV3(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			//1.找前驱节点
			predecessor := root.Left
			if predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			//2.判断前驱结点是否有右子树
			if predecessor.Right == nil {
				//无右子树,则 1.将右指针指向root 2.根节点左移
				predecessor.Right = root
				root = root.Left
			} else {
				//有右子树则 1.节点加入结果 2.节点右移
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return
}
