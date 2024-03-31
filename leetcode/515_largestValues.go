package leetcode

import "math"

/*
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	deq := []*TreeNode{}
	deq = append(deq, root)
	for len(deq) > 0 {
		size := len(deq)
		tempMax := math.MinInt64
		for i := 0; i < size; i++ {
			node := deq[0]
			deq = deq[1:]
			if node.Val > tempMax {
				tempMax = node.Val
			}
			if node.Left != nil {
				deq = append(deq, node.Left)
			}
			if node.Right != nil {
				deq = append(deq, node.Right)
			}
		}
		res = append(res, tempMax)
	}
	return res
}
