package Week_02

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	list := []int{root.Val}

	for _, n := range root.Children {
		list = append(list, preorder(n)...)
	}

	return list
}
