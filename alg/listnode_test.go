package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

//合并两有序链表
func mergeListNode(h1, h2 *ListNode) *ListNode {
	mergeListNode := &ListNode{}
	result := mergeListNode
	for h1 != nil && h2 != nil {
		if h1.Val > h2.Val {
			mergeListNode.Next = h2
			h2 = h2.Next
		} else {
			mergeListNode.Next = h1
			h1 = h1.Next
		}
		mergeListNode = mergeListNode.Next
	}

	if h1 != nil {
		mergeListNode.Next = h1
	}
	if h2 != nil {
		mergeListNode.Next = h2
	}
	return result.Next
}


