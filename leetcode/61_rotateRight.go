package leetcode

/*
题目：旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

题解：

*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	len := 0
	for ; p != nil; p = p.Next {
		len++
	}
	k %= len
	fast, slow := head, head
	for ; k != 0; k-- {
		fast = fast.Next
	}
	for ; fast.Next != nil; fast = fast.Next {
		slow = slow.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head

}
