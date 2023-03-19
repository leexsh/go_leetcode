package leetcode

/*
题目：二叉树的右视图
// 二叉树层次遍历
*/

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	que := []*TreeNode{}
	que = append(que, root)
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if i == size-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return res
}
