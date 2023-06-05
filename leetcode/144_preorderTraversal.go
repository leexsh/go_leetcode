package leetcode

/*
题目：二叉树的前序遍历
*/

// 递归
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

// 非递归
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
