package leetcode

import "math"

/*
题目：验证二叉搜索树
	给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
题解：
*/

// 递归
func isValidBST(root *TreeNode) bool {
	var helper func(*TreeNode, int, int) bool
	helper = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}
		if node.Val <= left || node.Val >= right {
			return false
		}
		return helper(node.Left, left, node.Val) && helper(node.Right, node.Val, right)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

func isValidBST1(root *TreeNode) bool {
	stack := []*TreeNode{}
	value := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if value > root.Val {
			return false
		}
		value = root.Val
		root = root.Right
	}
	return true
}
