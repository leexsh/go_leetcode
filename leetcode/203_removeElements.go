package leetcode

/*
题目： 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur, node := pre, head
	for node != nil {
		if node.Val == val {
			cur.Next = node.Next
			node = node.Next
		} else {
			cur = node
			node = node.Next
		}
	}
	return pre.Next
}
