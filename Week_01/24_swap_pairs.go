package Week_01

func swapPairs(head *ListNode) *ListNode {
	flag := 0
	sentinel := &ListNode{
		Val:  0,
		Next: head,
	}
	current := sentinel

	for current.Next != nil && current.Next.Next != nil {
		if flag%2 == 0 {
			next := current.Next
			nnext := current.Next.Next
			current.Next = nnext
			next.Next = nnext.Next
			nnext.Next = next
		}
		current = current.Next
		flag++
	}

	return sentinel.Next
}

func swapPairs2(head *ListNode) *ListNode {
	sentinel := &ListNode{
		0, head,
	}
	current := sentinel

	for current.Next != nil && current.Next.Next != nil {
		next := current.Next
		nnext := current.Next.Next
		current.Next = nnext
		next.Next = nnext.Next
		nnext.Next = next
		current = current.Next.Next
	}

	return sentinel.Next
}
