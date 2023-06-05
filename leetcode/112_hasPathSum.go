package leetcode

/*
题目：路径总和
给你二叉树的根节点root 和一个表示目标和的整数targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。
如果存在，返回 true ；否则，返回 false 。

题解：

*/
// 非递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	que := []*TreeNode{}
	que = append(que, root)
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node.Right == nil && node.Left == nil && node.Val == targetSum {
			return true
		}
		if node.Left != nil {
			node.Left.Val += node.Val
			que = append(que, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			que = append(que, node.Right)
		}
	}
	return false
}

// 递归
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum2(root.Left, targetSum-root.Val) || hasPathSum2(root.Right, targetSum-root.Val)
}
