package No111minDepth

import "fmt"

/**
 * Definition for a binary tree node.
 *
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	min := 0
	for len(queue) > 0 {
		min += 1
		length := len(queue)
		fmt.Println(queue)
		for length > 0 {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left == nil && queue[0].Right == nil {
				return min
			}
			queue = queue[1:]
			length--
		}
	}

	return min

}
