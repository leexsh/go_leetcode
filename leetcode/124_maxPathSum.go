package leetcode

import "math"

/*
题目：二叉树中最大路径之和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

*/

func maxPathSum(root *TreeNode) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := math.MinInt64

	var deal func(root *TreeNode) int
	deal = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, deal(root.Left))
		right := max(0, deal(root.Right))
		res = max(res, root.Val+left+right)
		return max(left, right) + root.Val
	}

	deal(root)
	return res
}
