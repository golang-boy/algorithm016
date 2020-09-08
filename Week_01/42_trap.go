package Week_01

//动态规划去找出i列的maxLeft及maxRight
func trap(height []int) int {
	size := len(height)
	if size < 3 {
		return 0
	}

	maxLefts := make([]int, size, size)
	maxRight := make([]int, size, size)

	maxLefts[0] = height[0]
	for i := 1; i < size; i++ {
		maxLefts[i] = maxInt(maxLefts[i-1], height[i])
	}

	maxRight[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = maxInt(maxRight[i+1], height[i])
	}

	capacity := 0
	for i := 1; i < size-1; i++ {
		capacity += minInt(maxLefts[i], maxRight[i]) - height[i]
	}

	return capacity
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func minInt(left, right int) int {
	if left <= right {
		return left
	}
	return right
}
