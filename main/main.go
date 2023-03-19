package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		if root != nil {
			root = root.Right
		}
	}

	return res
}
func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	l1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	l2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Right = l1
	l1.Left = l2
	fmt.Println(inorderTraversal(root))
}
