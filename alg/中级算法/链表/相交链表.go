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

//方法1：哈希集合
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	vis := map[*ListNode]bool{}

	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

//双指针
func getIntersectionNode2(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
