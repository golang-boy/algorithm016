package Week_03

func majorityElement(nums []int) int {

	m := map[int]int{}
	max := 0
	var maxElement  int

	length := len(nums) / 2

	for _, v :=range nums {
		m[v]++
		if m[v] > max {
			max = m[v]
			maxElement = v
		}
	}
	if max > length {
		return maxElement
	}
	return maxElement
}