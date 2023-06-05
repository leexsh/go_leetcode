package leetcode

/*
题目：路径总和2
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var dfs func(node *TreeNode, path []int, sum int)

	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && node.Val == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		dfs(node.Left, path, sum-node.Val)
		dfs(node.Right, path, sum-node.Val)
		path = path[:len(path)-1]
	}
	dfs(root, []int{}, targetSum)
	return res
}
