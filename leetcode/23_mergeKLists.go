package leetcode

// 23. 合并K个升序链表
/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
/*
https://leetcode.cn/problems/merge-k-sorted-lists/
*/
// 1. 暴力
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	var merge func(l1, l2 *ListNode) *ListNode
	merge = func(l1, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		if l1.Val < l2.Val {
			l1.Next = merge(l1.Next, l2)
			return l1
		} else {
			l2.Next = merge(l1, l2.Next)
			return l2
		}
	}
	pre := lists[0]
	for i := 1; i < len(lists); i++ {
		pre = merge(pre, lists[i])
	}
	return pre
}

// 2.分治
func mergeTwoList(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: -1}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return pre.Next
}
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return mergeTwoList(left, right)
}
