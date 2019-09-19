package main

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}
	isNeedReverse := false
	count := 0
	deep := 0
	h := &ListNode{}
	nodes := make([]*ListNode, 0)
	for {
		count += 1
		deep += 1
		if count == 1 {
			h = head
		}
		if count < k {
			if head.Next == nil {
				nodes = append(nodes, h)
				break
			}
			head = head.Next
		}
		if count == k {
			c := head.Next
			head.Next = nil
			nodes = append(nodes, h)
			if c == nil {
				isNeedReverse = true
				break
			}
			head = c
			count = 0
		}
	}
	if deep == k {
		return  reverseList(nodes[0])
	}
	if deep < k {
		return nodes[0]
	}
	eLast := &ListNode{}
	result := &ListNode{}
	isLast := false
	for key, node := range nodes {
		// 最后一个判断node长度决定是够反转，每一个node都要遍历并取得最后一个
		if key < len(nodes) - 1 {
			newNode := reverseList(node)
			if key == 0 {
				result = newNode
			}
			if isLast {
				eLast.Next = newNode
				isLast = false
			}
			for {
				if newNode.Next == nil {
					eLast = newNode
					isLast = true
					break
				}
				newNode = newNode.Next
			}
			continue
		}
		if isNeedReverse && key == len(nodes) - 1 {
			lastNode := reverseList(node)
			eLast.Next = lastNode
		} else {
			eLast.Next = node
		}
	}
	return result
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}