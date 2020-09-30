package Week_03

import "sort"

// 给定一个可包含重复数字的序列，返回所有不重复的全排列。

type Status struct {
	nums  []int
	visit map[int]bool
	path  []int
	ans   [][]int
}

func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)
	s := &Status{nums: nums, visit: make(map[int]bool)}
	backtrack3(s)
	return s.ans
}

func backtrack3(s *Status) {
	if len(s.path) == len(s.nums) {
		s.ans = append(s.ans, append([]int{}, s.path...))
		return
	}

	for i := 0; i < len(s.nums); i++ {
		if s.visit[i] {
			continue
		}

		if i > 0 && !s.visit[i-1] && s.nums[i] == s.nums[i-1] {
			continue
		}
		s.path = append(s.path, s.nums[i])
		s.visit[i] = true
		backtrack3(s)
		s.path = s.path[:len(s.path)-1]
		s.visit[i] = false
	}
}
