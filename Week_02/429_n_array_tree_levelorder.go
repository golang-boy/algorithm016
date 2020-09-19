package Week_02

type Node struct {
	Val int
	Children []*Node
}

ans := [][]int{}

func levelOrder(root *Node) [][]int {

	if root == nil {
		return nil
	}

	m := map[int][]int{}
	m[0] = []int{root.Val}

	bfs(root, 1, m)

	ans := [][]int{}

	for i:=0; i<len(m);i++{
		ans = append(ans, m[i])
	}

	return ans
}

func bfs(root *Node, level int, m map[int][]int) {
	if root == nil {
		return
	}

	for _, n := range root.Children {
		m[level] = append(m[level], n.Val)
	    bfs(n, level+1, m)
	}
}