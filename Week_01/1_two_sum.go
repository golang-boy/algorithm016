package Week_01

//1. 两数之和
func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nil
	}

	dict := make(map[int]int, numsLen*2)
	for index, value := range nums {

		if pos, ok := dict[target-value]; ok {
			return []int{index, pos}
		}

		dict[value] = index
	}

	return nil
}
