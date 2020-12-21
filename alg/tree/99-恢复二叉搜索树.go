package main

func recoverTree(root *TreeNode) {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapped(nums)
	recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}

//迭代法
func recoverTreeV2(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

//Morris中序遍历
func recoverTreeV3(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 如果没有左孩子，则直接访问右孩子
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
