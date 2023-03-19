package leetcode

/*
题目： 填充每个节点的下一个右侧节点指针
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。

描述：
*/

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
	Next  *TNode
}

func connect(root *TNode) *TNode {
	if root == nil {
		return root
	}
	que := []*TNode{}
	que = append(que, root)
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if i+1 == size {
				node.Next = nil
			} else {
				node.Next = que[0]
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return root
}
