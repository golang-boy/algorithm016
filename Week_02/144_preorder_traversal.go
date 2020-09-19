package Week_02

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var ans = []int{}
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	ans = append(ans, root.Val)
	ans = append(ans, left...)
	return append(ans, right...)
}