package leetcode

/*
题目：翻转链表II
给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

题解：
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur := pre
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}
	node := cur.Next
	for i := 0; i < right-left; i++ {
		next := node.Next
		node.Next = next.Next
		next.Next = cur.Next
		cur.Next = next
	}
	return pre.Next
}
