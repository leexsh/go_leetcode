package leetcode

/*
题目： 二叉树的最大深度
*/
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

// 非递归-层次遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deq := []*TreeNode{}
	deq = append(deq, root)
	res := 0
	for len(deq) > 0 {
		size := len(deq)
		res++
		for i := 0; i < size; i++ {
			node := deq[0]
			deq = deq[1:]
			if node.Right != nil {
				deq = append(deq, node.Right)
			}
			if node.Left != nil {
				deq = append(deq, node.Left)
			}
		}
	}
	return res
}
