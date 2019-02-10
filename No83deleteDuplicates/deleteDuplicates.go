package No83deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{Val: head.Val, Next: nil}
	val := head.Val
	last := result
	for ; head != nil; {
		if head.Val == val {
			head = head.Next
			continue
		} else {
			val = head.Val

			newNode := &ListNode{Val: val, Next: nil}
			last.Next = newNode
			last = last.Next
			head = head.Next
			continue
		}
	}
	return result
}
