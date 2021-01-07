package main

import (
	"fmt"
)

//练手

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

//全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)

	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return curr
}

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}
