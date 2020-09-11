package Week_01

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	meetNode := getMeetNode(head)

	if meetNode == nil {
		return nil
	}

	/**
	关键证明点：当pos1位于环h位置，环有c个节点，pos2位于0位置，pos1 step=2,pos2 step =1,在c-h的位置pos1与pos2重合
	设pos1=h,pos2=0,进过x次移动，pos1与pos2重合
	h+2x = x+kc(k为自然数)
	当k=0时，h+x=0 可得h,x都为0才能满足
	当k=1事，为普通情况第一次相遇，h+2x = x+c即得x=c-h
	*/
	start := head
	for start != meetNode {
		start, meetNode = start.Next, meetNode.Next
	}

	return start

}

func getMeetNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return fast
		}
	}

	return nil
}
