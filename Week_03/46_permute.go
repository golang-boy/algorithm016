package Week_03

type statusNode struct {
	ans   [][]int
	nums  []int
	path  []int
	visit map[int]bool
}

func permute(nums []int) [][]int {
	s := &statusNode{nums: nums, path: []int{}, ans: [][]int{}, visit: make(map[int]bool)}
	backtrack2(s)
	return s.ans
}

func backtrack2(s *statusNode) {
	if len(s.path) == len(s.nums) {
		s.ans = append(s.ans, append([]int{}, s.path...))
		return
	}

	for i := 0; i < len(s.nums); i++ {
		if s.visit[s.nums[i]] {
			continue
		}

		s.path = append(s.path, s.nums[i])
		s.visit[s.nums[i]] = true
		backtrack2(s)
		s.path = s.path[:len(s.path)-1]
		s.visit[s.nums[i]] = false
	}
}
