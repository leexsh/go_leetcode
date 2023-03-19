package leetcode

/*
题目：分割链表

题解：

*/
func partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{
		Val:  -1,
		Next: nil,
	}, &ListNode{
		Val:  -1,
		Next: nil,
	}
	p1, p2 := l1, l2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}

	p1.Next = l2.Next
	return l1.Next

}
