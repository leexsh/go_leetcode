package interview

/*
from 滴滴面试
	1.两个有序数组合并去重
	2.链表排序
*/

func mergeTwoArray(nums1, nums2 []int) []int {
	l1, l2, l := len(nums1)-1, len(nums2)-1, 0
	res := make([]int, l1+l2)
	i, j := 0, 0
	for i <= l1 && j <= l2 {
		if nums1[i] < nums2[j] {
			res[l] = nums1[i]
			l++
			i++
			for i <= l1 && l > 0 {
				if nums1[i] == res[l-1] {
					i++
				} else {
					break
				}
			}
		} else if nums1[i] == nums2[j] {
			res[l] = nums1[i]
			i++
			j++
			l++
			for i <= l1 && l > 0 {
				if nums1[i] == res[l-1] {
					i++
				} else {
					break
				}
			}
			for j <= l2 && l > 0 {
				if nums2[j] == res[l-1] {
					j++
				} else {
					break
				}
			}
		} else {
			res[l] = nums2[j]
			j++
			l++
			for j <= l2 && l > 0 {
				if nums2[j] == res[l-1] {
					j++
				} else {
					break
				}
			}
		}
	}

	for i <= l1 {
		res[l] = nums1[i]
		l++
		i++
		for i <= l1 && l > 0 {
			if nums1[i] == res[l-1] {
				i++
			} else {
				break
			}
		}
	}
	for j <= l2 {
		res[l] = nums2[j]
		l++
		j++
		for j <= l2 && l > 0 {
			if nums2[j] == res[l-1] {
				j++
			} else {
				break
			}
		}
	}
	return res[:l]
}

// 2.链表排序
type ListNode struct {
	val  int
	next *ListNode
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pre := &ListNode{
		val:  -1,
		next: nil,
	}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			cur.next = l1
			l1 = l1.next
		} else {
			cur.next = l2
			l2 = l2.next
		}
		cur = cur.next
	}
	if l1 != nil {
		cur.next = l1
	}
	if l2 != nil {
		cur.next = l2
	}
	return pre.next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	fast, slow, mid := head, head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		mid = slow
		slow = slow.next
	}
	mid.next = nil
	// need to merge
	return mergeList(sortList(head), sortList(slow))
}
