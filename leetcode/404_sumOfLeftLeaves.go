package leetcode

/*
给定二叉树的根节点 root ，返回所有左叶子之和。

*/

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var isLeafNode func(root *TreeNode) bool
	isLeafNode = func(node *TreeNode) bool {
		return node.Left == nil && node.Right == nil
	}
	var dfs func(node *TreeNode) (res int)
	dfs = func(node *TreeNode) (res int) {
		if node.Left != nil && isLeafNode(node.Left) {
			res += node.Left.Val
		} else if node.Left != nil && !isLeafNode(node.Left) {
			res += dfs(node.Left)
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			res += dfs(node.Right)
		}
		return
	}
	return dfs(root)
}
