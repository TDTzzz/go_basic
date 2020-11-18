package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快速排序
func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	// [l ... r) 左闭右开的链表区间
	quick_sort(head, nil)
	return head
}

// [l ... r) 左闭右开的链表区间
func quick_sort(l, r *ListNode) {
	// 区间没有元素或者只有一个元素
	if l == r || l.Next == r {
		return
	}
	mid := getIndex(l, r)

	quick_sort(l, mid)
	quick_sort(mid.Next, r)
}

// 获取基准元素的坐标
func getIndex(l, r *ListNode) *ListNode {
	if l == r || l.Next == r {
		return l
	}
	// 头指针作为基准元素
	tmp := l.Val
	p1 := l
	p2 := l.Next
	// *定义两个辅助指针p1，p2,这两个指针均往next方向移动，移动的过程中保持p1之前的值都小于选定的pivot，
	//    p1和p2之间的值都大于选定的pivot，那么当p2走到末尾时交换p1的值与pivot便完成了一次切分
	for p2 != r {
		if p2.Val < tmp {
			// 交换 p2的值到 p1
			p1 = p1.Next
			p1.Val, p2.Val = p2.Val, p1.Val
		}
		p2 = p2.Next
	}
	// p1表示的是小于tmp的最后一个元素
	p1.Val, l.Val = l.Val, p1.Val
	return p1
}
