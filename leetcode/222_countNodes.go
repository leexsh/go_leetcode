package leetcode

/*
题目：完全二叉树的节点个数
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
*/

// 递归
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 迭代
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	que := []*TreeNode{}
	res := 0
	que = append(que, root)
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		res++
		if node.Right != nil {
			que = append(que, node.Right)
		}
		if node.Left != nil {
			que = append(que, node.Left)
		}
	}
	return res
}
