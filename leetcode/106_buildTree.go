package leetcode

/*
题目：中旭和后续构造二叉树
*/
// todo
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}
	root.Left = buildTree2(inorder[:i], postorder[:i])
	root.Right = buildTree2(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
