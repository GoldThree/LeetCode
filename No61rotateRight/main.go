package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	tempHead := head

	count := 0
	last := &ListNode{}
	for {
		count += 1
		if head.Next != nil {
			head = head.Next
			continue
		}
		last = head
		head = tempHead
		break
	}
	if k == count {
		return head
	}

	if count < k {
		if k%count == 0 {
			return head
		}
		k = k % count
	}

	num := 0
	for {
		num += 1
		if head.Next == nil {
			break
		}
		if num-1 == count-k-1 {
			currentNode := head
			newHead := currentNode.Next
			last.Next = tempHead
			currentNode.Next = nil
			return newHead
		}
		head = head.Next
	}

	return head
}
