package Week_03

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[0:i])+1], inorder[0:i])
	root.Right = buildTree(preorder[len(inorder[0:i])+1:], inorder[i+1:])
	return root
}
