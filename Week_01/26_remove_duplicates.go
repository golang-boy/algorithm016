package Week_01

//26. 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	slow, fast := 1, 1
	for fast < numsLen {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}
