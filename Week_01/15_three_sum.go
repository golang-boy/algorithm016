package Week_01

import "sort"

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0, 10)

	for i := 0; i < numsLen-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		x := i + 1
		y := numsLen - 1
		target := -nums[i]
		for x < y {
			if x != i+1 && nums[x-1] == nums[x] {
				x++
				continue
			}

			if y != numsLen-1 && nums[y] == nums[y+1] {
				y--
				continue
			}

			if target == nums[x]+nums[y] {
				result = append(result, []int{nums[i], nums[x], nums[y]})
				x++
			} else if nums[x]+nums[y] > target {
				y--
			} else {
				x++
			}
		}
	}

	return result
}
