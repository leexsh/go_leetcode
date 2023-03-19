package leetcode

/*
题目：删除链表中的重复元素
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
题解：

*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur, L, R := pre, pre, pre
	for cur.Next != nil {
		L = cur.Next
		R = cur.Next
		for R.Next != nil && L.Val == R.Next.Val {
			R = R.Next
		}
		if L == R {
			cur = cur.Next
		} else {
			cur.Next = R.Next
		}
	}
	return pre.Next
}
