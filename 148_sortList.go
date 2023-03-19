package leetcode

/*
题目：对链表进行排序

*/
// 插入排序超出时间限制
// 归并
func mergeTList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pre := &ListNode{
		Val:  -1,
		Next: nil,
	}
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
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast, slow, mid := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = slow
		slow = slow.Next
	}
	mid.Next = nil
	return mergeTList(sortList(head), sortList(slow))
}
