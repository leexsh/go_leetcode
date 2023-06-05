package leetcode

/*
题目：对链表进行插入排序
题解：https://leetcode.cn/problems/insertion-sort-list/solution/dui-lian-biao-jin-xing-cha-ru-pai-xu-by-leetcode-s/
*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{
		Val:  -1,
		Next: head,
	}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			node := pre
			for node.Next.Val <= cur.Val {
				node = node.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = node.Next
			node.Next = cur
		}
		cur = lastSorted.Next
	}
	return pre.Next
}
