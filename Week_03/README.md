学习笔记

* 多数元素 https://leetcode-cn.com/problems/majority-element/
* 电话号码的字母组合 https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
* N皇后问题 https://leetcode-cn.com/problems/n-queens/
* 二叉树最近公共祖先 https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
* 从前序与中序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
* 组合 https://leetcode-cn.com/problems/combinations/
* 全排列   https://leetcode-cn.com/problems/permutations/
* 全排列 II https://leetcode-cn.com/problems/permutations-ii/


# 总结
排列组合问题是决策树遍历问题，因此通过画图的方式构造状态树，形成一个n叉树，然后遍历该树

```go
type Status struct {
	nums  []int
	path  []int
	ans   [][]int
}

func backtrack3(s *Status) {
    // 终止判断
	if len(s.path) == len(s.nums) {
	}

	for i := 0; i < len(s.nums); i++ {
        // 进行剪枝
        // dosomething
		backtrack3(s)
        // undosomething

	}
}
```