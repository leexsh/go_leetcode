package leetcode

/*
题目：对称的二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

// 递归
func isSymmetric(root *TreeNode) bool {
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}
	return dfs(root, root)
}

// 队列
func isSymmetric1(root *TreeNode) bool {
	deq := []*TreeNode{}
	deq = append(deq, root)
	deq = append(deq, root)
	for len(deq) > 0 {
		p1 := deq[0]
		p2 := deq[1]
		deq = deq[2:]
		if p1 == nil && p2 == nil {
			continue
		}
		if p1 == nil || p2 == nil {
			return false
		}
		if p1.Val != p2.Val {
			return false
		}
		deq = append(deq, p1.Left)
		deq = append(deq, p2.Right)
		deq = append(deq, p1.Right)
		deq = append(deq, p2.Left)
	}
	return true
}
