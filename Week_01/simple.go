package Week_01

import (
	"log"
)

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

//21. 合并两个有序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}

		preHead = preHead.Next
	}

	if l1 != nil {
		preHead.Next = l1
	} else {
		preHead.Next = l2
	}

	return result.Next

}

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

//66. 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}

	return append([]int{1}, digits...)
}
