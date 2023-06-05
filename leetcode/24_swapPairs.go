package leetcode

// 24. 两两交换链表中的节点/*
//
// 题目：https://leetcode.cn/problems/swap-nodes-in-pairs/
// 题解： https://leetcode.cn/problems/swap-nodes-in-pairs/solution/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-91/
// */

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 顺序遍历
func swapPairs1(head *ListNode) *ListNode {
	pre := &ListNode{Val: -1, Next: head}
	temp := pre
	for temp.Next != nil && temp.Next.Next != nil {
		n1, n2 := temp.Next, temp.Next.Next
		temp.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		temp = n1
	}
	return pre.Next
}
