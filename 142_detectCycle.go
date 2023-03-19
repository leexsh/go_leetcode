package leetcode

/*
题目： 环形链表
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/
func detectCycle1(head *ListNode) *ListNode {
	var hasCycle func(head *ListNode) bool
	fast, slow := head, head
	hasCycle = func(head *ListNode) bool {
		if head == nil {
			return false
		}
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		}
		return false
	}
	b := hasCycle(head)
	if !b {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
