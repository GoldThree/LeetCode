package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	count := 1
	nodes := make([]*ListNode, 0)
	for {
		if head == nil || head.Next == nil {
			nodes = append(nodes, head)
			break
		}
		if (count+1)%2 != 0 {
			head = head.Next
			count += 1
			continue
		}
		tail := head
		h := head.Next
		head = head.Next.Next
		count += 2



		h.Next = tail
		tail.Next = nil
		nodes = append(nodes, h)
	}

	result := &ListNode{}
	for key, node := range nodes {
		if key == 0 {
			result = node
		}
		if key > 0 {
			nodes[key-1].Next.Next = node
		}
	}
	return result
}
