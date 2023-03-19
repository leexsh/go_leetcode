package leetcode

/*
题目：二叉树的层次遍历
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	deq := []*TreeNode{}
	deq = append(deq, root)
	for len(deq) > 0 {
		size := len(deq)
		temp := []int{}
		for i := 0; i < size; i++ {
			node := deq[0]
			deq = deq[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				deq = append(deq, node.Left)
			}
			if node.Right != nil {
				deq = append(deq, node.Right)
			}
		}
		t1 := make([]int, len(temp))
		copy(t1, temp)
		res = append(res, t1)
	}
	return res
}
