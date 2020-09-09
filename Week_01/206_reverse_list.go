package Week_01

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next, pre = pre, current
		current = next
	}
	return pre
}
