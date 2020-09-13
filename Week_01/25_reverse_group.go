package Week_01

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{0, head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = nex
	}

	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	pre := tail.Next
	for pre != tail {
		nex := p.Next
		p.Next = pre
		pre = p
		p = nex
	}

	return tail, head
}
