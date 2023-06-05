package leetcode

/*
题目：拷贝随机节点
	给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
	构造这个链表的深拷贝。深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
	新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
	例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
	返回复制链表的头节点。

*/

// 利用map
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	mp := make(map[*Node]*Node)
	cur := head
	for ; cur != nil; cur = cur.Next {
		mp[cur] = &Node{
			Val:    cur.Val,
			Next:   nil,
			Random: nil,
		}
	}
	cur = head
	for ; cur != nil; cur = cur.Next {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
	}
	return mp[head]
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{
			Val:    cur.Val,
			Next:   next,
			Random: nil,
		}
		cur = next
	}
	cur = head
	for cur != nil {
		next := cur.Next.Next
		cur.Next.Random = cur.Random.Next
		cur = next
	}

	cur = head
	res := head.Next
	for cur != nil {
		next := cur.Next.Next
		copy := cur.Next
		cur.Next = next
		if next.Next != nil {
			copy.Next = next.Next
		}
		cur = next
	}
	return res
}
