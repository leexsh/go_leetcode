package leetcode

/*
题目：有序链表转二叉搜索树


*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var preSlow *ListNode = nil
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{
		Val:   slow.Val,
		Left:  nil,
		Right: sortedListToBST(slow.Next),
	}
	if preSlow != nil {
		preSlow.Next = nil
		root.Left = sortedListToBST(head)
	}
	return root
}
