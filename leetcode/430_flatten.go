package leetcode

/*
你会得到一个双链表，其中包含的节点有一个下一个指针、一个前一个指针和一个额外的 子指针 。
这个子指针可能指向一个单独的双向链表，也包含这些特殊的节点。这些子列表可以有一个或多个自己的子列表，
以此类推，以生成如下面的示例所示的 多层数据结构 。
输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]
*/

type Node430 struct {
	Val   int
	Prev  *Node430
	Next  *Node430
	Child *Node430
}

/*
	有rand 走rand
	没有就走next
	利用stack回退
*/

func flatten_430(root *Node430) *Node430 {
	if root == nil {
		return nil
	}
	stack := []*Node430{}
	cur := root
	for cur != nil {
		if cur.Child != nil {
			if cur.Next != nil {
				stack = append(stack, cur.Next)
			}
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
		}
		if cur.Next != nil {
			cur = cur.Next
			continue
		}
		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Prev = cur
		cur.Next = node
		cur = node
	}
	return root
}
