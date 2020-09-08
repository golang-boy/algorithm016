package Week_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
