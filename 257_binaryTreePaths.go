package leetcode

import "strconv"

/*
题目：二叉树的所有路径
*/
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode, s string)
	dfs = func(root *TreeNode, s string) {
		if root != nil {
			s += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				res = append(res, s)
			} else {
				s += "->"
				dfs(root.Left, s)
				dfs(root.Right, s)
			}
		}
	}
	dfs(root, "")
	return res
}
