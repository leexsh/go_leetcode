package leetcode

/*
题目：翻转链表
*/
func reverseList1(head *ListNode) *ListNode {
	var cur *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}
	return cur

}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
