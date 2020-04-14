package No445addTwoNumbersII

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Nums := getNums(l1)
	l2Nums := getNums(l2)
	l1Length := len(l1Nums)
	l2Length := len(l2Nums)

	maxLength := max(l1Length, l2Length)

	nextNode := &ListNode{}
	nextNode = nil
	curNode := &ListNode{}

	extal := 0
	for i := 0; i < maxLength; i++ {

		n1 := 0
		if i < l1Length {
			n1 = l1Nums[i]
		}

		n2 := 0
		if i < l2Length {
			n2 = l2Nums[i]
		}

		value := n1 + n2 + extal
		extal = 0
		if value >= 10 {
			extal = value / 10
			value = value % 10
		}

		curNode = &ListNode{}
		curNode.Val = value
		curNode.Next = nextNode
		nextNode = curNode
	}

	if extal != 0 {
		curNode = &ListNode{}
		curNode.Val = extal
		curNode.Next = nextNode
	}

	return curNode
}

func getNums(l *ListNode) []int {

	nums := make([]int, 0)
	if l == nil {
		return nums
	}

	for node := l; node != nil; node = node.Next {
		nums = append([]int{node.Val}, nums...)
	}

	return nums
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// func getNum(l *ListNode) int {

//     result := 0
//     if l == nil {
//         return 0
//     }

//     for node :=l;node != nil;node=node.Next {
//         fmt.Println(node.Val, result)
//         result = result*10 +node.Val
//     }

//     return result
// }

func getList(num int) *ListNode {

	if num == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	nextNode := &ListNode{}
	nextNode = nil
	curNode := &ListNode{}

	for {
		if num <= 0 {
			break
		}
		curNode = &ListNode{}
		value := num % 10
		curNode.Val = value
		curNode.Next = nextNode
		nextNode = curNode
		num = num / 10
	}

	return curNode

}
