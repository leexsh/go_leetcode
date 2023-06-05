package leetcode

/*
题目：二叉树转为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

*/

// 中序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	list := []*TreeNode{}
	stack := []*TreeNode{}
	for len(stack) > 0 && root != nil {
		for root.Right != nil {
			list = append(list, root)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}
