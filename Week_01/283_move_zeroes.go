package Week_01

//283. 移动零
func moveZeroes(nums []int) {
	//0的位置
	left, right := -1, -1

	for index, value := range nums {
		if value == 0 {
			right = index
			if left < 0 {
				left = index
			}
		} else if left >= 0 {
			nums[index], nums[left] = 0, value
			left++
			right++
		}
	}
}
