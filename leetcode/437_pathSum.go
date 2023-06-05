package leetcode

/*
题目： 路径总和 III
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
*/
func pathSum3(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		cnt := 0
		if root.Val == sum {
			cnt++
		}
		cnt += dfs(root.Left, sum-root.Val)
		cnt += dfs(root.Right, sum-root.Val)
		return cnt
	}
	return dfs(root, targetSum) + dfs(root.Right, targetSum) + dfs(root.Left, targetSum)
}
