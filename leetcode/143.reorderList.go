package leetcode

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

// 1.中間节点
// 2.逆序
// 3.合并两个链表
func midNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseNode(head *ListNode) *ListNode {
	var cur *ListNode = nil
	for head != nil {
		node := head.Next
		head.Next = cur
		cur = head
		head = node
	}
	return cur
}

func mergeCur(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := midNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseNode(l2)
	mergeCur(l1, l2)
}
