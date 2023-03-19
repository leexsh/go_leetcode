package leetcode

//  删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{Val: -1}
	pre.Next = head
	var fast, slow *ListNode = head, pre
	for ; n != 0; n-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
