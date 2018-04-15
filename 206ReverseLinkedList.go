
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//0
	if head == nil {
		return head
	}
	var cur1 *ListNode
	cur1 = head
	var cur2 *ListNode
	if (*cur1).Next != nil {
		//2
		cur2 = cur1
		cur1 = (*cur1).Next

		if (*cur1).Next == nil {
			(*cur1).Next = cur2
			(*cur2).Next = nil
			return cur1
		}
		cur1 = head
		cur2 = nil
		for (*cur1).Next != nil {
			cur3 := (*cur1).Next
			(*cur1).Next = cur2
			cur2 = cur1
			cur1 = cur3

		}
		(*cur1).Next = cur2

	} else {
		//1
		return cur1
	}
	return cur1
}