package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位值, 只可能为0或1
	promotion := 0

	// 结果表的头结点
	var head *ListNode
	// 保存结果表的尾结点
	var rear *ListNode
	for nil != l1 || nil != l2 {
		sum := 0
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += promotion
		promotion = 0

		if sum >= 10 {
			promotion = 1
			sum = sum % 10
		}

		node := &ListNode{
			sum,
			nil,
		}

		if nil == head {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = node
		}

	}

	if promotion > 0 {
		rear.Next = &ListNode{
			promotion,
			nil,
		}
	}

	return head
}
