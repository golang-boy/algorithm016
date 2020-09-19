package Week_02

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var ans = []int{}
	left := inorderTraversal(root.Left)
    ans = append(ans, left...)
    ans = append(ans, root.Val)
	right := inorderTraversal(root.Right)
	return append(ans, right...)
}