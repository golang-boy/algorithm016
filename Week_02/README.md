# 学习笔记

* 有效的字母异位词：https://leetcode-cn.com/problems/valid-anagram/description/
* 字母异位词分组：https://leetcode-cn.com/problems/group-anagrams/
* 二叉树的中序遍历：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
* 二叉树的前序遍历：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
* N叉树的前序遍历：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
* N叉树的层序遍历：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
* 丑数：https://leetcode-cn.com/problems/chou-shu-lcof/
* 前 K 个高频元素：https://leetcode-cn.com/problems/top-k-frequent-elements/


## 总结

* N叉树前后序遍历模板

```go
	if root == nil {
		return []int{}
	}

    // root.Val放着是前序 
	// list := []int{root.Val}

	for _, n := range root.Children {
		list = append(list, preorder(n) ... )
	}

    // root.Val放着是后序 
	// list := append(list, root.Val)
	return list
```

* 二叉树前中后序遍历模板
```go

	if root == nil {
		return []int{}
	}

    var ans = []int{}

	left := preorderTraversal(root.Left)
    right := preorderTraversal(root.Right)
    
    // 前序  <
    ans = append(ans, root.Val)
    ans = append(ans, left...)
	ans = append(ans, right...)

	// 中序  A
    ans = append(ans, left...)
    ans = append(ans, root.Val)
	ans = append(ans, right...)

	// 后序 >
    ans = append(ans, left...)
	ans = append(ans, right...)
    ans = append(ans, root.Val)
	return ans
```

* 树结构迭代遍历时
  * 深度优先搜索：用栈结构缓存中间结果
  * 广度优先搜索：通过队列结构缓存中间结果

* 丑数的关键是确定下一个丑数
    通过有序队列减少重复计算，提升性能
    * 将计算过的丑数放入有序队列中
    * 确定下一个丑数后，将其移出有序队列，放入新计算的丑数
    * 该有序队列可以通过大顶堆数据结构来实现

* 涉及基本数据结构操作 (增删改查)
   * 数组
   * 堆
   * 树
