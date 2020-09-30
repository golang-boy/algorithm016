package Week_03

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//  输入: n = 4, k = 2
//  输出:
//  [
//    [2,4],
//    [3,4],
//    [2,3],
//    [1,2],
//    [1,3],
//    [1,4],
//  ]
//

type StrategyNode struct {
	n      int
	k      int
	nums   []int //路径
	result [][]int
}

func combine(n int, k int) [][]int {
	s := &StrategyNode{k: k, n: n, nums: []int{}, result: [][]int{}}
	backtrack1(s, 1)
	return s.result
}

func backtrack1(s *StrategyNode, begin int) {
	// 结束条件
	if len(s.nums) == s.k {
		var res []int
		res = append(res, s.nums...)
		s.result = append(s.result, res)
		return
	}

	for i := begin; i <= s.n; i++ {
		s.nums = append(s.nums, i)
		backtrack1(s, i+1)
		s.nums = s.nums[:len(s.nums)-1]
	}
}
