package leetcode

/*
题目： 二叉树的锯齿形遍历
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	flag := true
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
		if flag {
			t1 := make([]int, len(temp))
			copy(t1, temp)
			res = append(res, t1)
		} else {
			t1 := make([]int, 0, len(temp))
			for i := len(temp) - 1; i >= 0; i-- {
				t1 = append(t1, temp[i])
			}
			res = append(res, t1)
		}

		flag = !flag
	}
	return res
}
