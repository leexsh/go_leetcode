package leetcode

/*
题目：回文链表
https://leetcode.cn/problems/palindrome-linked-list/solution/shou-hua-tu-jie-hui-wen-lian-biao-kuai-man-zhi-zhe/
*/
func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	var pre *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	var head2 *ListNode = nil
	pre.Next = nil
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}
