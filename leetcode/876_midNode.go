package leetcode

/*
题目：链表的中间节点

*/
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
