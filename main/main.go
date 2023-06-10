package main

import (
	"fmt"
	"leexsh/leetcode/leetcode"
)

func sumOfLeftLeaves(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	var isLeafNode func(root *leetcode.TreeNode) bool
	isLeafNode = func(node *leetcode.TreeNode) bool {
		return node.Left == nil && node.Right == nil
	}
	var dfs func(node *leetcode.TreeNode) (res int)
	dfs = func(node *leetcode.TreeNode) (res int) {
		if node.Left != nil && isLeafNode(node.Left) {
			res += node.Left.Val
		} else if node.Left != nil && !isLeafNode(node) {
			res += dfs(node.Left)
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			res += dfs(node.Right)
		}
		return
	}
	return dfs(root)
}

func main() {
	root := &leetcode.TreeNode{
		Val: 3,
		Left: &leetcode.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &leetcode.TreeNode{
			20,
			&leetcode.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			&leetcode.TreeNode{
				7, nil, nil,
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(root))
}
