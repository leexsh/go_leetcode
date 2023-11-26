package leetcode

/*
K 个一组翻转链表
題目：https://leetcode.cn/problems/reverse-nodes-in-k-group/
题解：https://leetcode.cn/problems/reverse-nodes-in-k-group/solution/by-aeon27-qye1/
https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/248591/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
*/

// todo
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Val: -1,
		Next: head,
	}
	pre := newHead
	for head != nil {
		tail := pre
		for i := 0 ; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHead.Next
			}
		}
		nex := tail.Next
		head, tail := reverseList(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return newHead.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}
