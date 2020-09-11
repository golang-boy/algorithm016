package Week_01

import "testing"

func TestDetectCycle(t *testing.T) {
	head := &ListNode{
		Val:  3,
		Next: nil,
	}

	start := head

	start.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	start = start.Next
	start.Next = &ListNode{
		Val:  0,
		Next: nil,
	}
	start = start.Next
	start.Next = &ListNode{
		Val:  -4,
		Next: head.Next,
	}

	detectCycle(head)

}
