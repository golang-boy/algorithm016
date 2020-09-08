package Week_01

//88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1

	m -= 1
	n -= 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
		pos--
	}

	if n >= 0 {
		for i := 0; i <= n; i++ {
			nums1[i] = nums2[i]
		}
	}
}
