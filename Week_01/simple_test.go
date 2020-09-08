package Week_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 6
	rotate1(nums, k)
	assert.Equal(t, nums, []int{4, 5, 6, 7, 8, 9, 1, 2, 3})
}

func TestRotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 6
	rotate2(nums, k)
	assert.Equal(t, nums, []int{4, 5, 6, 7, 8, 9, 1, 2, 3})
}

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	result := []int{1, 1, 2, 3, 4, 4}

	resultList := mergeTwoLists(l1, l2)
	for _, value := range result {
		assert.Equal(t, value, resultList.Val)
		resultList = resultList.Next
	}

	assert.Nil(t, resultList)
}
