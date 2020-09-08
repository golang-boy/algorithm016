package Week_01

//283. 移动零
func moveZeroes1(nums []int) {
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

func moveZeroes2(nums []int) {
	//需要交换的位置
	pos := 0

	for index, value := range nums {
		if value != 0 {
			nums[index], nums[pos] = nums[pos], nums[index]
			pos++
		}
	}
}

func moveZeroes3(nums []int) {
	//0的个数
	count := 0

	for index, value := range nums {
		if value == 0 {
			count++
		} else if count > 0 {
			nums[index-count], nums[index] = nums[index], 0
		}
	}
}
