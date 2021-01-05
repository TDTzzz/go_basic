package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func reverseListV2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

func reverseListV3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else {
		var dfs func(head, p *ListNode)
		var newhead = head
		var p = head.Next
		for newhead.Next != nil {
			newhead = newhead.Next
		}
		dfs = func(head, p *ListNode) {
			if p == nil {
				return
			}
			dfs(head.Next, p.Next)
			p.Next = head
			head.Next = nil
		}
		dfs(head, p)
		return newhead
	}
}

func reverseListV4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
