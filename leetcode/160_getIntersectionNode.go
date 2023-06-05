package leetcode

/*
题目：相交链表
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}
	lenA, lenB := 0, 0
	cur := headA
	for cur != nil {
		lenA++
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		lenB++
		cur = cur.Next
	}
	if lenA > lenB {
		val := lenA - lenB
		for val > 0 {
			headA = headA.Next
			val--
		}
	} else {
		val := lenB - lenA
		for val > 0 {
			val--
			headB = headB.Next
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

// 牛逼
// https://leetcode.cn/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
