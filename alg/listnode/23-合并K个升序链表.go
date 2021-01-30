package listnode

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func mergeKLists(lists []*ListNode) *ListNode {
	h := HeapMin{}
	heap.Init(&h)
	node := &ListNode{}
	cur := node
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, lists[i])
	}

	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		cur.Next = min
		cur = cur.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return node.Next
}

type HeapMin []*ListNode

func (h HeapMin) Len() int {
	return len(h)
}

func (h HeapMin) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *HeapMin) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *HeapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func (h *HeapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
