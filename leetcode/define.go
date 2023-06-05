package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
