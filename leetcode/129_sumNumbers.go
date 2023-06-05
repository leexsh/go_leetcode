package leetcode

/*
题目：求根节点到叶节点数字之和
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

*/

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	que := []*TreeNode{}
	res := 0
	que = append(que, root)
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if node.Left == nil && node.Right == nil {
				res += node.Val
			}
			if node.Left != nil {
				node.Left.Val = node.Val*10 + node.Left.Val
				que = append(que, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*10 + node.Right.Val
				que = append(que, node.Right)
			}
		}
	}
	return res
}
