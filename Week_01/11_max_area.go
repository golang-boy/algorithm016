package Week_01

func maxArea(height []int) int {
	count := len(height)
	if count < 2 {
		return 0
	}

	start, end := 0, count-1
	maxValue := 0
	minHeight := 0
	for start < end {
		if height[start] < height[end] {
			minHeight = height[start]
			start++
		} else {
			minHeight = height[end]
			end--
		}

		maxValue = maxInt(maxValue, minHeight*(end-start+1))

		for minHeight >= height[start] && start < end {
			start++
		}

		for minHeight >= height[end] && start < end {
			end--
		}
	}

	return maxValue
}

func maxArea2(height []int) (area int) {
	count := len(height)
	if count < 2 {
		return
	}

	start, end := 0, count-1
	for start < end {
		area = maxInt(area, (end-start)*minInt(height[start], height[end]))
		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}

	return
}
