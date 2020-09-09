package Week_01

import "log"

//189. 旋转数组
func rotate1(nums []int, k int) {
	numsLen := len(nums)
	if numsLen < 1 {
		return
	}

	k = k % numsLen
	if k == 0 {
		return
	}

	count := 0
	for start := 0; count < numsLen; start++ {
		/**
		分3种情况：1，numsLen能被k整除，2，不能被整除,但start不需要增加，3，不能被整除，start需要增加
		情况1：start需要不断增加到k-1

		情况2：需要证明当count==numsLen-1，current一直不等于start,循环过程中没有访问重复元素，即访问了所有元素，
		相当于从current出发，步长为k，第一次访问重复元素位置count=x,第二次count=y,当count=y-x时current==start与条件矛盾

		情况3：当current==start但count<numLens,新的start=start+1,需要证明新的start并没有被访问过，且新的循环元素也没有被原有循环遍历过
		当start+1没有被访问过，那么start+1起的循环元素也不会与原来的重复，不然start+1也应该重复
		当start+1已经被访问过，count=x,那么从start经过2x能访问start+2,经过3x能访问start+3，由情况2可知，一个循环中，元素不会重复，可知从start无重复访问所有元素
		*/
		log.Println("------------------", start)
		current := (start + k) % numsLen
		preValue := nums[start]

		for start != current {
			preValue, nums[current] = nums[current], preValue
			current = (current + k) % numsLen
			count++
		}
		nums[start] = preValue
		count++
	}
}

func rotate2(nums []int, k int) {
	numsLen := len(nums)
	if numsLen < 1 {
		return
	}

	k = k % numsLen
	if k == 0 {
		return
	}

	reverse(nums, 0, numsLen)
	reverse(nums, 0, k)
	reverse(nums, k, numsLen)
}

func reverse(nums []int, start int, end int) {
	i := start
	j := end - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
